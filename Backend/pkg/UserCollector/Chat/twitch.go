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
	"time"
)

const ConfigName = "Twitch.json"

type ready struct {
	ready   bool
	changed bool
}

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

	OAuth OAuth `json:"oauth"`
	// Caches Ready data
	ready
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

	if tc.ApiKey == "" {
		return false
	}

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
		err := tc.saveConfig()
		if err != nil {
			log.Println(err)
		}
	}

	tc.client = twitch.NewClient(tc.ChannelSender, tc.ApiKey)
	tc.client.Capabilities = []string{twitch.TagsCapability}

	tc.ready.changed = true
	tc.Ready()
	go tc.getIds()

	r.Get("/", tc.get)

	r.Post("/", tc.post)
	r.Get("/ready/", tc.isReady)
	r.Get("/oauth/", tc.getOAuth)
	r.Get("/std-oauth/", tc.setOAuthStandard)
	r.Get("/adv/", tc.getAdv)
	r.Post("/adv", tc.postAdv)

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

	if tc.OAuth.Color != 0 && tc.OAuth.ready {
		err := tc.announce(msg)
		if err != nil {
			log.Println(err)
		}
		return
	}

	tc.Lock()
	defer tc.Unlock()

	tc.client.OnConnect(func() {
		tc.client.Say(tc.Channel, msg)
		time.Sleep(50 * time.Millisecond)
		tc.client.Disconnect()
	})
	_ = tc.client.Connect()
}

// CollectGuesses collect the user votes until the context cancels
func (tc *TwitchChat) CollectGuesses(ctx context.Context, collect func(id int, user, message string)) {
	tc.Lock()
	defer tc.Unlock()
	client := tc.client

	client.Join(tc.Channel)

	if tc.OAuth.Color != 0 && tc.OAuth.ready {
		client.OnConnect(func() {})
		err := tc.announce(tc.StartMsg)
		if err != nil {
			log.Println(err)
		}
	} else {
		client.OnConnect(func() {
			client.Say(tc.Channel, tc.StartMsg)
		})
	}

	tc.client.OnPrivateMessage(func(m twitch.PrivateMessage) {
		if !strings.HasPrefix(m.Message, tc.Prefix) {
			return
		}

		id, err := strconv.Atoi(m.User.ID)
		if err != nil {
			log.Printf("Error parsing %s user id (%s): %v", m.User.Name, m.User.ID, err)
			return
		}

		collect(id, strings.ToLower(m.User.Name), strings.TrimSpace(strings.TrimLeft(m.Message, tc.Prefix)))
	})

	ctx, cancel := context.WithCancel(ctx)
	go func() {
		err := client.Connect()
		if err != nil && !errors.Is(err, twitch.ErrClientDisconnected) {
			log.Println(err)
			cancel()
		}
	}()

	if c := ctx.Done(); c != nil {
		<-c
	} else {
		log.Println(ctx.Err())
	}

	if tc.OAuth.Color != 0 && tc.OAuth.ready {
		err := tc.announce(tc.EndMsg)
		if err != nil {
			log.Println(err)
		}
	} else {
		client.Say(tc.Channel, tc.EndMsg)
	}

	// make sure, a message is sent before the client is closed
	time.Sleep(50 * time.Millisecond)
	client.Disconnect()
}

// formatFinalMessage replaces $RESULT and $WINNERS for the message for the chat
func (tc *TwitchChat) formatFinalMessage(winners []string, result float64) (res string) {
	res = strings.Replace(tc.FinalMsg, "$RESULT", fmt.Sprintf("%.2f", result), -1)
	if len(winners) == 0 {
		winners = []string{"keiner"}
	}
	res = strings.Replace(res, "$WINNERS", strings.Join(winners, ", "), -1)

	return
}

// saveConfig saves the config into the configPath file
func (tc *TwitchChat) saveConfig() error {
	file, err := os.OpenFile(tc.configPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(tc, "", "  ")
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	return err
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

	var client struct {
		Channel       string `json:"channel"`
		Prefix        string `json:"prefix"`
		ApiKey        string `json:"apiKey"`
		ChannelSender string `json:"channelSender"`
		MsgBegin      string `json:"msgBegin"`
		MsgEnd        string `json:"msgEnd"`
		MsgFinal      string `json:"msgFinal"`
	}
	err := ctx.BodyParser(&client)
	if err != nil {
		return err
	}

	sender, key := tc.ChannelSender, tc.ApiKey

	tc.Channel = client.Channel
	tc.Prefix = client.Prefix
	tc.ApiKey = client.ApiKey
	tc.ChannelSender = client.ChannelSender
	tc.StartMsg = client.MsgBegin
	tc.FinalMsg = client.MsgFinal
	tc.EndMsg = client.MsgEnd

	tc.checkEmptySender()

	// recreate the client if credentials change
	if sender != tc.ChannelSender || key != tc.ApiKey {
		tc.ready.changed = true
		tc.client = twitch.NewClient(tc.ChannelSender, tc.ApiKey)
		tc.client.Capabilities = []string{twitch.TagsCapability}
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
