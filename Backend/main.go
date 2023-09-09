package main

import (
	"Killspiel/pkg/Killspiel"
	"Killspiel/pkg/helper"
	"embed"
	"github.com/goccy/go-json"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.ConfigDefault))
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/", websocket.New(func(conn *websocket.Conn) {

	}))

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
