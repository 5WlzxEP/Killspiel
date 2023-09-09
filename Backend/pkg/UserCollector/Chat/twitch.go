package Chat

import (
	"Killspiel/pkg/config"
	"context"
	"errors"
	"fmt"
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
)

const ConfigName = "Twitch.json"

type TwitchChat struct {
	client        *twitch.Client
	Channel       string `json:"channel"`
	Prefix        string `json:"prefix"`
	ApiKey        string `json:"apiKey"`
	ChannelSender string `json:"channelSender"`
	StartMsg      string `json:"msgBegin"`
	EndMsg        string `json:"msgEnd"`
	FinalMsg      string `json:"msgFinal"`
	configPath    string
	// Caches Ready data
	ready struct {
		ready   bool
		changed bool
	}
	sync.Mutex
}

func (tc *TwitchChat) Ready() bool {
	if !tc.ready.changed {
		return tc.ready.ready
	}
	tc.Lock()
	defer tc.Unlock()

	tc.ready.changed = false
	tc.ready.ready = false

	var err2 error
	tc.client.OnConnect(func() {
		err2 = tc.client.Disconnect()
	})
	err := tc.client.Connect()
	if err != nil && !errors.Is(err, twitch.ErrClientDisconnected) {
		return false
	}
	if err2 != nil {
		return false
	}

	tc.ready.ready = true
	return true
}

func New(configPath string, r fiber.Router) (*TwitchChat, string) {
	tc := newConfig(configPath)
	if tc == nil {
		tc = &TwitchChat{
			configPath: path.Join(configPath, ConfigName),
		}
		tc.saveConfig()
	}

	tc.client = twitch.NewClient(tc.ChannelSender, tc.ApiKey)

	tc.ready.changed = true
	tc.Ready()

	r.Get("/", tc.get)

	r.Post("/", tc.post)
	r.Get("/ready/", tc.isReady)

	return tc, "Twitchchat"
}

func newConfig(configPath string) *TwitchChat {
	conf := path.Join(configPath, ConfigName)

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

	tc.checkEmptySender()

	return &tc
}

func (tc *TwitchChat) AnnounceResult(winners []string, correctGuess float64) {
	msg := tc.formatFinalMessage(winners, correctGuess)

	tc.Lock()
	defer tc.Unlock()

	// TODO Remove block
	err := tc.client.Connect()
	if err != nil {
		return
	}
	tc.client.Say(tc.Channel, msg)
	tc.client.Disconnect()
}

// CollectGuesses collect the user votes until the context cancels
func (tc *TwitchChat) CollectGuesses(ctx context.Context, collect func(id int, user, message string)) {
	tc.Lock()
	defer tc.Unlock()
	client := tc.client

	client.Join(tc.Channel)

	client.OnConnect(func() {
		client.Say(tc.Channel, tc.StartMsg)
	})

	tc.client.OnPrivateMessage(func(m twitch.PrivateMessage) {
		if !strings.HasPrefix(m.Message, tc.Prefix) {
			return
		}

		id, err := strconv.Atoi(m.User.ID)
		if err != nil {
			log.Printf("Error parsing %s user id (%s): %v", m.User.Name, m.User.ID, err)
		}

		collect(id, strings.ToLower(m.User.Name), strings.TrimLeft(m.Message, tc.Prefix))
	})

	ctx, cancel := context.WithCancel(ctx)
	go func() {
		err := client.Connect()
		if err != nil {
			cancel()
		}
	}()

	if c := ctx.Done(); c != nil {
		<-c
	}

	client.Say(tc.Channel, tc.EndMsg)
	client.Disconnect()

}

// formatFinalMessage replaces $RESULT and $WINNERS for the message for the chat
func (tc *TwitchChat) formatFinalMessage(winners []string, result float64) (res string) {
	res = strings.Replace(tc.FinalMsg, "$RESULT", fmt.Sprintf("%.2f", result), -1)
	res = strings.Replace(tc.FinalMsg, "$WINNERS", strings.Join(winners, ", "), -1)
	return
}

// saveConfig saves the config into the configPath file
func (tc *TwitchChat) saveConfig() error {
	file, err := os.OpenFile(tc.configPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(tc)
}

// get returns the current config json encoded
func (tc *TwitchChat) get(ctx *fiber.Ctx) error {
	return ctx.JSON(tc)
}

// post "gets" new Config from the frontend, checks ist and saves it.
//
// might recreate the client if the client credentials change.
//
// returns a 409 if CollectGuesses runs because it needs the TwitchChat.client running
func (tc *TwitchChat) post(ctx *fiber.Ctx) error {
	success := tc.TryLock()
	if !success {
		return ctx.Status(http.StatusConflict).SendString("Killspiel laeuft gerade, keine aenderungen moeglich")
	}
	defer tc.Unlock()

	client, auth := tc.ChannelSender, tc.ApiKey
	err := ctx.BodyParser(tc)
	if err != nil {
		return err
	}

	tc.checkEmptySender()

	// recreate the client if credentials change
	if client != tc.ChannelSender || auth != tc.ApiKey {
		tc.ready.changed = true
		tc.client = twitch.NewClient(tc.ChannelSender, tc.ApiKey)
	}

	err = tc.saveConfig()
	if err != nil {
		return err
	}

	ctx.Status(http.StatusNoContent)
	return nil
}

// checkEmptySender check is sender is empty, if so sets channel as sender
func (tc *TwitchChat) checkEmptySender() {
	// check if channel sender and channel are the same and there ChannelSender empty
	if tc.ChannelSender == "" {
		tc.ChannelSender = tc.Channel
	}
}

func (tc *TwitchChat) isReady(ctx *fiber.Ctx) error {
	ctx.Status(http.StatusOK)
	if tc.Ready() {
		return ctx.SendString("true")
	}
	return ctx.SendString("false")
}
