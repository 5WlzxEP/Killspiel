package RiotApi

import (
	"Killspiel/pkg/config"
	"context"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"net/http"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
	"time"
)

const configName = "RiotApi.json"

var (
	servers = []string{"br1", "eun1", "euw1", "jp1", "kr", "la1", "la2", "na1", "oc1", "ph2", "ru", "sg2", "th2", "tr1", "tw2", "vn2"}
	//regions = []string{"americas", "asia", "europe", "sea"}
)

type General struct {
	Intervall time.Duration `json:"intervall"`
}

type Api struct {
	LoL             `json:"lol"`
	General         `json:"general"`
	client          *fasthttp.Client
	ready           bool
	summoner        *Summoner
	currentSummoner *Summoner
	currentGame     *CurrentGameInfo
	configPath      string
}

func New(configPath string, r fiber.Router) (*Api, string) {
	api := getApi(configPath)

	r.Get("/", api.get)
	r.Post("/", api.post)
	r.Post("/lol/", api.postLoL)
	r.Get("/ready/", api.getReady)

	return api, "RiotApi"
}

func getApi(cp string) *Api {
	configPath := path.Join(cp, configName)

	var api = new(Api)
	var err error
	if config.ExistsAndFile(configPath) {
		var f *os.File
		f, err = os.OpenFile(configPath, os.O_RDONLY, 0)
		if err == nil {
			err = json.NewDecoder(f).Decode(api)
		}
	} else {
		createFile(configPath, api)
	}

	if err != nil {
		logger.Printf("An error occured trying to load RiotApi: %v", err)
	}

	api.client = &fasthttp.Client{
		NoDefaultUserAgentHeader:      true,
		Dial:                          (&fasthttp.TCPDialer{DNSCacheDuration: time.Hour}).Dial,
		DisableHeaderNamesNormalizing: true,
	}

	api.configPath = configPath
	api.isServerOrDefault()
	api.region = getRegion(api.Server)
	api.ready = api.initReady()

	return api
}

func createFile(configPath string, api *Api) {
	file, err := os.Create(configPath)
	if err != nil {
		return
	}
	defer file.Close()
	data, err := json.MarshalIndent(api, "", "  ")
	if err != nil {
		return
	}
	_, _ = file.Write(data)
}

func (a *Api) initReady() bool {
	if !isServer(a.Server) {
		return false
	}
	if !validKey(a.ApiKey, a.client) {
		return false
	}
	summoner, err := getLoLSummonerByAccount(a.Name, a.Tag, a.region, a.Server, a.ApiKey, a.client)
	if err != nil || summoner == nil {
		return false
	}
	a.summoner = summoner
	return true
}

func (a *Api) Begin(ctx context.Context, cancelFunc context.CancelFunc, dbInfo chan<- string) {
	done := ctx.Done()

	for ; ; time.Sleep(a.Intervall) {
		select {
		case <-done:
			a.currentGame = nil
			return
		default:
			if a.checkInGame() {
				a.currentSummoner = a.summoner
				dbInfo <- fmt.Sprintf("LoL,%s#%s,%d,%s", a.Name, a.Tag, a.currentGame.GameId, a.LoL.Kategorie)
				cancelFunc()
				return
			}
		}
	}
}

func (a *Api) Prefix() string {
	return "LoL"
}

func (a *Api) Resolve(dbinfo string) (float64, error) {
	parts := strings.Split(dbinfo, ",")
	if len(parts) != 4 {
		return 0, errors.New("dbinfo not a valid LoL info")
	}

	nameTag := strings.Split(parts[1], "#")
	if len(nameTag) != 2 {
		return 0, errors.New("nameTag must contain a hashtag")
	}
	summoner, err := getLoLSummonerByAccount(nameTag[0], nameTag[1], a.region, a.Server, a.ApiKey, a.client)

	if err != nil {
		return 0, err
	}

	a.currentSummoner = summoner

	matchId, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		return 0, err
	}

	match, err := a.getMatchById(matchId)
	if err != nil {
		return 0, err
	}

	return a.getValueFloat(match, parts[3]), nil
}

func (a *Api) checkInGame() bool {
	game, err := a.getActiveGameById(a.summoner.Puuid)
	if err != nil || (a.currentGame != nil && game.GameId == a.currentGame.GameId) {
		return false
	}
	a.currentGame = game
	return true
}

func (a *Api) Result(ctx context.Context, c chan float64) {
	done := ctx.Done()

	// started manual must be stop manual
	if a.currentGame == nil {
		return
	}

	for ; ; time.Sleep(a.Intervall) {
		select {
		case <-done:
			return
		default:
			match, err := a.getMatchById(a.currentGame.GameId)
			if err != nil {
				continue
			}
			c <- a.getValueFloat(match, a.LoL.Kategorie)
			break
		}
	}
}

func (a *Api) save() {
	f, err := os.OpenFile(a.configPath, os.O_WRONLY|os.O_TRUNC, 0)
	if err != nil {
		return
	}
	defer f.Close()
	bytes, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return
	}

	_, err = f.Write(bytes)
	if err != nil {
		logger.Printf("An Error occurred writing %s: %v\n", configName, err)
	}
}

func getRegion(server string) string {
	switch server {
	case "na1", "br1", "la1", "la2":
		return "americas"
	case "kr", "jp1":
		return "asia"
	case "eun1", "euw1", "tr1", "ru":
		return "europe"
	case "oc1", "ph2", "sg2", "th2", "tw2", "vn2":
		return "sea"
	}
	return "europe"
}

func (a *Api) Ready() bool {
	return a.ready
}

func validKey(apiKey string, client *fasthttp.Client) bool {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://euw1.api.riotgames.com/lol/status/v4/platform-data")

	req.Header.Set("X-Riot-Token", apiKey)
	req.Header.SetMethod(http.MethodGet)
	res := fasthttp.AcquireResponse()

	err := client.Do(req, res)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)
	if err != nil {
		return false
	}

	return !(res.StatusCode() > 300 || res.StatusCode() < 200)
}

// isServerOrDefault checks Api.Server is a Server or sets it to the first default, which is a server or to euw1
func (a *Api) isServerOrDefault(default_ ...string) {
	if isServer(a.Server) {
		return
	}
	for _, d := range default_ {
		if isServer(d) {
			a.Server = d
			return
		}
	}
	a.Server = "euw1"
}

func isServer(server string) bool {
	return slices.Contains(servers, server)
}

func (a *Api) get(ctx *fiber.Ctx) error {
	if a.summoner != nil {
		a.LoL.ProfileIcon = a.summoner.ProfileIconId
	} else {
		a.LoL.ProfileIcon = -1
	}

	return ctx.JSON(a)
}

func (a *Api) post(ctx *fiber.Ctx) error {
	var gen General
	err := ctx.BodyParser(&gen)
	if err != nil {
		return err
	}

	a.Intervall = max(gen.Intervall*time.Second, 5*time.Second)
	a.save()

	return ctx.SendStatus(http.StatusNoContent)
}

func (a *Api) getReady(ctx *fiber.Ctx) error {

	ctx.Status(http.StatusOK)
	if a.Ready() {
		return ctx.SendString("true")
	}
	return ctx.SendString("false")
}
