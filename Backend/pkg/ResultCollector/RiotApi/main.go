package RiotApi

import (
	"Killspiel/pkg/config"
	"errors"
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
	Key    string
	Server string
	region string
	client *fasthttp.Client
	ready
	summoner *Summoner
}

var (
	servers = []string{"NA", "BR", "LAN", "LAS", "KR", "JP", "EUNE", "EUW", "TR", "RU", "OCE", "PH2", "SG2", "TH2", "TW2", "VN2"}
	regions = []string{"AMERICAS", "ASIA", "EUROPE", "SEA"}
)

func New(configPath string, r fiber.Router) (*Api, string) {
	api := getApi(configPath)

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

func (a Api) setRegion() {
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

func (a Api) Ready() bool {
	if !a.changed {
		return a.val
	}
	a.val = slices.Contains(servers, a.Server) &&
		slices.Contains(regions, a.region) &&
		a.ValidKey()

	return a.val
}

func (a Api) ValidKey() bool {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://euw1.api.riotgames.com/lol/status/v4/platform-data\n")

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
