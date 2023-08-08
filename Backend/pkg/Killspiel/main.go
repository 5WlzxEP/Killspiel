package Killspiel

import (
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

	router.CreateApiGroup(app)
}

func Run() {

}
