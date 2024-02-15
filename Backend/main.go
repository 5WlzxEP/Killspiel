package main

import (
	"Killspiel/pkg/Killspiel"
	"Killspiel/pkg/Websocket"
	"Killspiel/pkg/helper"
	"context"
	"embed"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	//go:embed frontend_build/*
	frontendBuild embed.FS
	Build         string
)

func main() {
	log.Printf("Version: %s\n", Build)

	app := fiber.New(fiber.Config{
		JSONDecoder: json.Unmarshal,
		JSONEncoder: json.Marshal,
		Network:     "tcp",
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

	ctx, cancel := context.WithCancel(context.Background())

	finished := make(chan struct{})

	go Killspiel.Run(ctx, finished)

	go func() {
		err := app.Listen(helper.EnvOrDefault("KILLSPIEL_HOST", ":8088"))
		if err != nil {
			panic(err)
		}
	}()

	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Kill, os.Interrupt)

	<-shutdown
	cancel()
	log.Println("Shutting down")
	err := app.ShutdownWithTimeout(time.Minute)
	if err != nil {
		log.Println(err)
	}
	<-finished
}
