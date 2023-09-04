package Killspiel

import (
	"Killspiel/pkg/Leaderboard"
	"Killspiel/pkg/User"
	"Killspiel/pkg/UserCollector"
	"Killspiel/pkg/config"
	"Killspiel/pkg/database"
	"Killspiel/pkg/router"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

func Init(app *fiber.App) {
	path, err := config.FindConfOrDefault()
	if err != nil {
		log.Fatal(err)
	}

	conf, err := config.GetConfig(path)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Init(path)
	if err != nil {
		log.Fatal(err)
	}

	api := router.CreateApiGroup(app)
	UserCollector.Init(path, conf.UserCollector, api.Group("/collector"))

	Leaderboard.Init(api.Group("/leaderboard"))
	User.Init(api.Group("/user"))

	// Redirecting user access per link, because /user/:id is routed through svelte
	app.Get("/user/:id/", func(ctx *fiber.Ctx) error {
		ctx.Status(http.StatusFound).Location(fmt.Sprintf("/Redirect/?%s", ctx.OriginalURL()))
		return nil
	})
}

func Run() {

}
