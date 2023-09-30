package Chat

import (
	"Killspiel/pkg/Websocket"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
	"strconv"
	"strings"
	"time"
)

var httpClient = &fasthttp.Client{}

type Msg struct {
	Message string `json:"message"`
	Color   string `json:"color"`
}

func (tc *TwitchChat) announce(msg string) error {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(fmt.Sprintf("https://api.twitch.tv/helix/chat/announcements?broadcaster_id=%d&moderator_id=%d", tc.OAuth.broadcasterId, tc.OAuth.moderatorId))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tc.OAuth.AccessToken))
	req.Header.Set("Client-Id", tc.OAuth.ClientId)
	req.Header.Set("content-type", "application/json")

	bytes, err := json.Marshal(Msg{
		Message: msg,
		Color:   codeToColor(tc.OAuth.Color),
	})
	if err != nil {
		return err
	}

	req.SetBodyRaw(bytes)

	err = httpClient.Do(req, nil)
	return err
}

type Users struct {
	Data []struct {
		Id              string    `json:"id"`
		Login           string    `json:"login"`
		DisplayName     string    `json:"display_name"`
		Type            string    `json:"type"`
		BroadcasterType string    `json:"broadcaster_type"`
		Description     string    `json:"description"`
		ProfileImageUrl string    `json:"profile_image_url"`
		OfflineImageUrl string    `json:"offline_image_url"`
		ViewCount       int       `json:"view_count"`
		CreatedAt       time.Time `json:"created_at"`
	} `json:"data"`
}

func (tc *TwitchChat) getIds() {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(fmt.Sprintf("https://api.twitch.tv/helix/users?login=%s&login=%s", tc.ChannelSender, tc.Channel))
	req.Header.SetMethod("POST")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tc.OAuth.AccessToken))
	req.Header.Set("Client-Id", tc.OAuth.ClientId)
	req.Header.Set("content-type", "application/json")

	resp := fasthttp.AcquireResponse()
	err := httpClient.Do(req, resp)
	if err != nil {
		return
	}

	if resp.StatusCode() != 200 {
		tc.OAuth.ready = false
		Websocket.Broadcast([]byte("Error: LoL: " + string(resp.Body())))
		return
	}

	var users Users
	err = json.Unmarshal(resp.Body(), &users)
	if err != nil {
		return
	}

	tc.OAuth.ready = true

	for _, user := range users.Data {
		if user.Login == strings.ToLower(tc.Channel) {
			id, err := strconv.Atoi(user.Id)
			if err != nil {
				tc.OAuth.ready = false
				continue
			}
			tc.OAuth.broadcasterId = id
		} else if user.Login == strings.ToLower(tc.ChannelSender) {
			id, err := strconv.Atoi(user.Id)
			if err != nil {
				tc.OAuth.ready = false
				continue
			}
			tc.OAuth.moderatorId = id
		}
	}

	if tc.OAuth.ready == false {
		return
	}

	// check if user is moderator
	req.Reset()

	req.SetRequestURI(fmt.Sprintf("https://api.twitch.tv/helix/moderation/blocked_terms?broadcaster_id=%d&moderator_id=%d", tc.OAuth.broadcasterId, tc.OAuth.moderatorId))
	req.Header.SetMethod("GET")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tc.OAuth.AccessToken))
	req.Header.Set("Client-Id", tc.OAuth.ClientId)

	resp.Reset()
	err = httpClient.Do(req, resp)
	if err != nil {
		tc.OAuth.ready = false
		return
	}
	if resp.Header.StatusCode() > 299 {
		tc.OAuth.ready = false
		Websocket.Broadcast([]byte("Error: LoL: " + string(resp.Body())))
	}

}

func codeToColor(color uint8) string {
	switch color {
	case 2:
		return "blue"
	case 3:
		return "green"
	case 4:
		return "orange"
	case 5:
		return "purple"
	default:
		return "primary"

	}
}
