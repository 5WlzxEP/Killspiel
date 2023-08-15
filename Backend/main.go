package main

import (
	"Killspiel/pkg/Killspiel"
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"net/http"
)

var (
	//go:embed frontend_build/*
	frontendBuild         embed.FS
	DisableStartupMessage = "false"
)

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: DisableStartupMessage == "true",
	})

	Killspiel.Init(app)

	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(frontendBuild),
		PathPrefix:   "frontend_build",
		MaxAge:       0,
		NotFoundFile: "r.html",
	}))

	go Killspiel.Run()

	panic(app.Listen(":8088"))
}
