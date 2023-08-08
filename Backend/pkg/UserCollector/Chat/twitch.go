package Chat

import (
	"context"
	"fmt"
	"github.com/gempir/go-twitch-irc/v4"
	"strconv"
	"strings"
)

type TwitchChat struct {
	client       *twitch.Client
	channel      string
	startMessage string
	endMessage   string
	prefix       string
}

func (tc *TwitchChat) Ready() error {
	//TODO implement me
	panic("implement me")
}

func (tc *TwitchChat) AnnounceResult(winners []string, correctGuess float64) {
	msg := tc.formatEndMessage(winners, correctGuess)

	tc.client.Connect()
	tc.client.Say(tc.channel, msg)
	tc.client.Disconnect()
}

func New(username, oauth, channel string) (*TwitchChat, error) {
	client := twitch.NewClient(username, oauth)

	chat := &TwitchChat{
		client:  client,
		channel: channel,
	}

	return chat, nil
}

func (tc *TwitchChat) CollectGuesses(ctx context.Context) map[string]float64 {
	client := tc.client

	client.Join(tc.channel)

	client.OnConnect(func() {
		client.Say(tc.channel, tc.startMessage)
	})

	guesses := make(map[string]float64)

	client.OnPrivateMessage(func(m twitch.PrivateMessage) {
		if !strings.HasPrefix(m.Message, tc.prefix) {
			return
		}

		guess, err := strconv.ParseFloat(strings.TrimLeft(m.Message, tc.prefix), 64)
		if err != nil {
			return
		}
		guesses[m.User.Name] = guess
	})

	go client.Connect()

	if c := ctx.Done(); c != nil {
		<-c
	}

	client.Say(tc.channel, tc.endMessage)
	client.Disconnect()

	return guesses
}

func (tc *TwitchChat) formatEndMessage(winners []string, result float64) (res string) {

	res = strings.Replace(tc.endMessage, "$RESULT", fmt.Sprintf("%.0f", result), -1)
	res = strings.Replace(tc.endMessage, "$Winners", strings.Join(winners, ", "), -1)
	return
}
