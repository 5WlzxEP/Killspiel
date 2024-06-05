package main

import (
	"Killspiel/pkg/Killspiel"
	"Killspiel/pkg/Websocket"
	"Killspiel/pkg/config"
	"cmp"
	"context"
	"embed"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	path2 "path"
	"time"
)

var (
	//go:embed frontend_build/*
	frontendBuild embed.FS
	Build         string = "Development"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: time.DateTime,
	})

	log.WithField("Version", Build).Println("Starting")

	app := fiber.New(fiber.Config{
		JSONDecoder: json.Unmarshal,
		JSONEncoder: json.Marshal,
		Network:     "tcp",
	})

	logFile := getLogFile()
	defer logFile.Close()

	app.Use(logger.New(logger.Config{
		Format:        "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
		TimeFormat:    "15:04:05",
		TimeZone:      "Local",
		TimeInterval:  500 * time.Millisecond,
		Output:        logFile,
		DisableColors: false,
	}))
	app.Use(cors.New(cors.ConfigDefault))
	Websocket.Init(app.Group("/ws"))

	Killspiel.Init(app)

	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(frontendBuild),
		PathPrefix:   "frontend_build",
		MaxAge:       60 * 60,
		NotFoundFile: "r.html",
	}))

	ctx, cancel := context.WithCancel(context.Background())

	finished := make(chan struct{})

	go Killspiel.Run(ctx, finished)

	go func() {
		err := app.Listen(cmp.Or(os.Getenv("KILLSPIEL_HOST"), ":8088"))
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

func getLogFile() *os.File {
	path, err := config.FindConfOrDefault()
	if err != nil {
		log.Error("Error finding config path", err)
		os.Exit(-1)
	}

	filepath := path2.Join(path, "webserver.log")

	if s, err := os.Stat(filepath); err != nil && s != nil {
		if s.Size() > 24*1024*1024 {
			newName := path2.Join(path, "webserver.log.1")
			_ = os.Remove(newName)

			err := os.Rename(filepath, newName)
			if err != nil {
				panic(err)
			}
		}
	}

	logFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return logFile
}
