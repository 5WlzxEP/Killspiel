package main

import (
	"Killspiel/pkg/Killspiel"
	"Killspiel/pkg/Websocket"
	"Killspiel/pkg/helper"
	"embed"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"net/http"
)

var (
	//go:embed frontend_build/*
	frontendBuild embed.FS
)

func main() {
	app := fiber.New(fiber.Config{
		JSONDecoder: json.Unmarshal,
		JSONEncoder: json.Marshal,
	})

	app.Use(logger.New())
	app.Use(cors.New(cors.ConfigDefault))
	Websocket.Init(app.Group("/ws"))

	Killspiel.Init(app)

	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(frontendBuild),
		PathPrefix:   "frontend_build",
		MaxAge:       30 * 60,
		NotFoundFile: "r.html",
	}))

	go Killspiel.Run()

	panic(app.Listen(helper.EnvOrDefault("KILLSPIEL_PORT", ":8088")))

}
