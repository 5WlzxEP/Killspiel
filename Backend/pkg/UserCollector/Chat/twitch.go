package Chat

import (
	"Killspiel/pkg/config"
	"context"
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
)

const CONFIG_NAME = "Twitch.json"

type TwitchChat struct {
	client        *twitch.Client
	Channel       string `json:"channel"`
	Prefix        string `json:"prefix"`
	ApiKey        string `json:"apiKey"`
	ChannelSender string `json:"channelSender"`
	StartMsg      string `json:"msgBegin"`
	EndMsg        string `json:"msgEnd"`
	FinalMsg      string `json:"msgFinal"`
	guesses       map[string]int
	configPath    string
	sync.Mutex
}

func (tc *TwitchChat) Ready() bool {
	err := tc.client.Connect()
	if err != nil {
		return false
	}
	err = tc.client.Disconnect()
	if err != nil {
		return false
	}
	return true
}

func (tc *TwitchChat) AnnounceResult(winners []string, correctGuess int) {
	msg := tc.formatFinalMessage(winners, correctGuess)

	tc.Lock()
	defer tc.Unlock()

	tc.client.Connect()
	tc.client.Say(tc.Channel, msg)
	tc.client.Disconnect()
}

func New(configPath string, r fiber.Router) *TwitchChat {
	chat := newConfig(configPath)
	if chat == nil {
		chat = &TwitchChat{
			guesses:    make(map[string]int),
			configPath: path.Join(configPath, CONFIG_NAME),
		}
		chat.saveConfig()
	}

	r.Get("/", chat.get)
	//r.Get("/", func(ctx *fiber.Ctx) error {
	//	return ctx.SendString("ABCDEFGH")
	//})
	r.Post("/", chat.post)

	return chat
}

func newConfig(configPath string) *TwitchChat {
	conf := path.Join(configPath, CONFIG_NAME)
	if !config.ExistsAndFile(conf) {
		return nil
	}
	f, err := os.Open(conf)
	if err != nil {
		return nil
	}
	defer f.Close()
	var tc TwitchChat
	err = json.NewDecoder(f).Decode(&tc)
	if err != nil {
		return nil
	}
	tc.configPath = conf
	tc.guesses = make(map[string]int)
	return &tc
}

func (tc *TwitchChat) CollectGuesses(ctx context.Context) map[string]int {
	tc.Lock()
	defer tc.Unlock()
	client := tc.client

	client.Join(tc.Channel)

	client.OnConnect(func() {
		client.Say(tc.Channel, tc.StartMsg)
	})

	clear(tc.guesses)

	// TODO move into client creation
	client.OnPrivateMessage(func(m twitch.PrivateMessage) {
		if !strings.HasPrefix(m.Message, tc.Prefix) {
			return
		}

		guess, err := strconv.Atoi(strings.TrimLeft(m.Message, tc.Prefix))
		if err != nil {
			return
		}
		tc.guesses[m.User.Name] = guess
	})

	go client.Connect()

	if c := ctx.Done(); c != nil {
		<-c
	}

	client.Say(tc.Channel, tc.EndMsg)
	client.Disconnect()

	return tc.guesses
}

func (tc *TwitchChat) formatFinalMessage(winners []string, result int) (res string) {
	res = strings.Replace(tc.FinalMsg, "$RESULT", strconv.Itoa(result), -1)
	res = strings.Replace(tc.FinalMsg, "$Winners", strings.Join(winners, ", "), -1)
	return
}

func (tc *TwitchChat) saveConfig() error {
	file, err := os.OpenFile(tc.configPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(tc)
}

func (tc *TwitchChat) get(ctx *fiber.Ctx) error {
	return ctx.JSON(tc)
}

// post "gets" new Config from the frontend, checks ist and saves it.
//
// Might recreate the client if the client credentials change.
//
// Returns a 409 if CollectGuesses runs, because it needs the TwitchChat.client running
func (tc *TwitchChat) post(ctx *fiber.Ctx) error {
	success := tc.TryLock()
	if !success {
		ctx.Status(http.StatusConflict)
		return nil
	}
	defer tc.Unlock()

	client, auth := tc.ChannelSender, tc.ApiKey
	err := ctx.BodyParser(tc)
	if err != nil {
		return err
	}

	// check if channel sender and channel are the same and there ChannelSender empty
	if tc.ChannelSender == "" {
		tc.ChannelSender = tc.Channel
	}

	// recreate the client if credentials change
	if client != tc.ChannelSender || auth != tc.ApiKey {
		tc.client = twitch.NewClient(tc.ChannelSender, tc.ApiKey)

	}

	err = tc.saveConfig()
	if err != nil {
		return err
	}

	ctx.Status(http.StatusOK)
	return nil
}
