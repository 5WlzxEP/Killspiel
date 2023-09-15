package RiotApi

import (
	"Killspiel/pkg/config"
	"context"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"os"
	"path"
	"slices"
	"time"
)

const configName = "RiotApi.json"

type ready struct {
	val, changed bool
}

type Api struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	// Intervall is in seconds
	Intervall    time.Duration `json:"interval"`
	SummonerName string
	region       string
	client       *fasthttp.Client
	ready
	summoner    *Summoner
	currentGame *CurrentGameInfo
}

func (a *Api) Begin(ctx context.Context, cancelFunc context.CancelFunc, dbInfo chan<- string) {
	done := ctx.Done()

	for ; ; time.Sleep(a.Intervall * time.Second) {
		select {
		case <-done:
			return
		default:
			if a.checkInGame() {
				dbInfo <- fmt.Sprintf("LoL,%s,%d", a.summoner.Name, a.currentGame.GameId)
				cancelFunc()
				return
			}
		}
	}
}

func (a *Api) checkInGame() bool {
	game, err := a.getActiveGameById(a.summoner.Id)
	if err != nil {
		return false
	}
	a.currentGame = game
	return true
}

func (a *Api) Result(ctx context.Context, c chan float64) {

}

var (
	servers = []string{"NA", "BR", "LAN", "LAS", "KR", "JP", "EUNE", "EUW", "TR", "RU", "OCE", "PH2", "SG2", "TH2", "TW2", "VN2"}
	regions = []string{"AMERICAS", "ASIA", "EUROPE", "SEA"}
)

func New(configPath string, r fiber.Router) (*Api, string) {
	api := getApi(configPath)

	r.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(404)
	})

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
		api.setRegion()
	} else {
		err = errors.New("no config yet")
	}

	if err != nil {
		log.Printf("An error occured trying to load RiotApi: %v", err)
	}

	api.client = &fasthttp.Client{
		NoDefaultUserAgentHeader:      true,
		Dial:                          (&fasthttp.TCPDialer{DNSCacheDuration: time.Hour}).Dial,
		DisableHeaderNamesNormalizing: true,
	}

	api.ready.changed = true

	return api
}

func (a *Api) setRegion() {
	switch a.Server {
	case "NA", "BR", "LAN", "LAS":
		a.region = "AMERICAS"
	case "KR", "JP":
		a.region = "ASIA"
	case "EUNE", "EUW", "TR", "RU":
		a.region = "EUROPE"
	case "OCE", "PH2", "SG2", "TH2", "TW2", "VN2":
		a.region = "SEA"
	}
}

func (a *Api) Ready() bool {
	if !a.changed {
		return a.val
	}
	a.val = slices.Contains(servers, a.Server) &&
		slices.Contains(regions, a.region) &&
		a.validKey() &&
		a.validSummoner()

	return a.val
}

func (a *Api) validKey() bool {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://euw1.api.riotgames.com/lol/status/v4/platform-data")

	req.Header.Set("X-Riot-Token", a.Key)
	req.Header.SetMethod(http.MethodGet)
	res := fasthttp.AcquireResponse()

	err := a.client.Do(req, res)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)
	if err != nil {
		return false
	}

	return !(res.StatusCode() > 300 || res.StatusCode() < 200)
}

func (a *Api) validSummoner() bool {
	summoner, err := a.getSummonerByName(a.SummonerName)
	if err != nil {
		return false
	}
	a.summoner = summoner
	return true
}
