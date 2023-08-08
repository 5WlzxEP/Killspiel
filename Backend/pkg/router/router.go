package router

import "github.com/gofiber/fiber/v2"

func CreateApiGroup(app *fiber.App) fiber.Router {
	api := app.Group("/api")

	//api.Use("/leaderboard", func() {})
	return api
}
