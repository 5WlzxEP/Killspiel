package UserCollector

import (
	"Killspiel/pkg/UserCollector/Chat"
	"Killspiel/pkg/config"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"slices"
	"time"
)

type Collectors []*cols

type cols struct {
	Name      string `json:"name"`
	Collector `json:"-"`
	Ready     bool `json:"ready"`
}

type Collector interface {
	Ready() bool
	CollectGuesses(ctx context.Context, collector func(id int, user, guess string))
	AnnounceResult(winners []string, correctGuess float64)
}

var collectors Collectors
var currentCollector *cols
var collectTime time.Duration
var conf *config.Config

func Init(configPath string, config *config.Config, r fiber.Router) {

	r.Get("/", get)
	r.Post("/", post)

	twitchChat, name := Chat.New(configPath, r.Group("/chat"))

	collectors = append(collectors, &cols{
		Name:      name,
		Collector: twitchChat,
	})

	collectTime = config.UserCollector.Duration * time.Second
	conf = config

	i := slices.IndexFunc(collectors, func(c *cols) bool {
		return c.Name == config.UserCollector.Collector
	})
	if i == -1 {
		return
	}

	if col := collectors[i]; col.Collector.Ready() {
		currentCollector = col
	}
}

func get(ctx *fiber.Ctx) error {
	for _, col := range collectors {
		col.Ready = col.Collector.Ready()
	}

	return ctx.JSON(map[string]any{
		"current": currentCollector,
		"all":     collectors,
		"time":    conf.UserCollector.Duration,
	})
}

func post(ctx *fiber.Ctx) error {
	var con config.UserCollect

	err := ctx.BodyParser(&con)
	if err != nil {
		return err
	}

	if con.Duration <= 1 {
		return ctx.Status(http.StatusBadRequest).SendString("Duration must be at least 1 second.")
	}

	index := slices.IndexFunc(collectors, func(c *cols) bool {
		return c.Name == con.Collector
	})

	var col *cols
	if index == -1 {
		goto noCollector
	}
	if col = collectors[index]; !col.Collector.Ready() {
		goto noCollector
	}

	conf.UserCollector.Duration = con.Duration
	conf.UserCollector.Collector = col.Name

	collectTime = con.Duration * time.Second
	currentCollector = col

	err = conf.Save()
	if err != nil {
		return err
	}

	return ctx.SendStatus(http.StatusNoContent)

noCollector:
	conf.UserCollector.Duration = con.Duration
	collectTime = con.Duration * time.Second

	err = conf.Save()
	if err != nil {
		log.Println(err)
		return err
	}
	return ctx.Status(299).SendString("Collector invalid, not saved but rest saved")
}
