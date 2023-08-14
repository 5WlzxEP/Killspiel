package Killspiel

import (
	"Killspiel/pkg/UserCollector"
	"Killspiel/pkg/config"
	"Killspiel/pkg/database"
	"Killspiel/pkg/router"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Init(app *fiber.App) {
	path, err := config.FindConfigPath()
	if err != nil {
		log.Fatal(err)
	}

	err = database.Init(path)
	if err != nil {
		log.Fatal(err)
	}

	api := router.CreateApiGroup(app)
	UserCollector.Init(path, api.Group("/collector"))
}

func Run() {

}
