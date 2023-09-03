package main

import (
	"Killspiel/pkg/Killspiel"
	"Killspiel/pkg/helper"
	"embed"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"net/http"
)

var (
	//go:embed frontend_build/*
	frontendBuild         embed.FS
	DisableStartupMessage = "false"
)

func main() {
	//goland:noinspection GoBoolExpressions
	app := fiber.New(fiber.Config{
		DisableStartupMessage: DisableStartupMessage == "true",
		JSONDecoder:           json.Unmarshal,
		JSONEncoder:           json.Marshal,
	})

	app.Use(logger.New())

	Killspiel.Init(app)

	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(frontendBuild),
		PathPrefix:   "frontend_build",
		MaxAge:       30 * 60,
		NotFoundFile: "r.html",
	}))

	go Killspiel.Run()

	panic(app.Listen(helper.EnvOrDefault("Address", ":8088")))

}
