package Chat

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type auth struct {
	Scope            string `query:"scope"`
	State            string `query:"state"`
	AccessToken      string `query:"access_token"`
	TokenType        string `query:"token_type"`
	Error            string `query:"error"`
	ErrorDescription string `query:"error_description"`
}

type OAuth struct {
	AccessToken   string `json:"apiKey"`
	ClientId      string `json:"clientId"`
	Color         uint8  `json:"color"`
	broadcasterId int
	moderatorId   int
	ready         bool
}

func (tc *TwitchChat) getAdv(ctx *fiber.Ctx) error {
	return ctx.JSON(tc.OAuth)
}

func (tc *TwitchChat) postAdv(ctx *fiber.Ctx) error {
	var oauth struct {
		AccessToken string `json:"apiKey"`
		ClientId    string `json:"clientId"`
		Color       uint8  `json:"color"`
	}

	if err := ctx.BodyParser(&oauth); err != nil {
		return err
	}

	if oauth.AccessToken != "" {
		tc.OAuth.AccessToken = oauth.AccessToken
	}

	if oauth.ClientId != "" {
		tc.OAuth.ClientId = oauth.ClientId
	}

	tc.OAuth.Color = oauth.Color

	_ = tc.saveConfig()
	go tc.getIds()

	return ctx.SendStatus(http.StatusNoContent)
}

func (tc *TwitchChat) getOAuth(ctx *fiber.Ctx) error {
	var a auth

	if err := ctx.QueryParser(&a); err != nil {
		return err
	}

	if a.Error != "" {
		return errors.New(a.Error + ": " + a.ErrorDescription)
	}

	tc.OAuth.AccessToken = a.AccessToken

	_ = tc.saveConfig()
	go tc.getIds()

	return ctx.Redirect("/Redirect/?/settings/collector/twitchchat/")
}

func (tc *TwitchChat) setOAuthStandard(ctx *fiber.Ctx) error {
	tc.OAuth.ClientId = "6yqxsptsgw5bffjno9b1q9dtr2nlf8"

	return ctx.Redirect("https://id.twitch.tv/oauth2/authorize?response_type=token&client_id=6yqxsptsgw5bffjno9b1q9dtr2nlf8&redirect_uri=http://localhost:8088/oauth/&scope=chat:read+chat:edit+channel:moderate+whispers:read+whispers:edit+moderator:manage:announcements+moderation:read+moderator:read:blocked_terms", http.StatusTemporaryRedirect)
}
