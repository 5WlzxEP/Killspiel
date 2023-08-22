package UserCollector

import (
	"Killspiel/pkg/UserCollector/Chat"
	"Killspiel/pkg/config"
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"slices"
)

type Collectors []cols

type cols struct {
	Name string `json:"name"`
	Collector
}

type Collector interface {
	Ready() bool
	CollectGuesses(ctx context.Context, collector func(id int, user, guess string))
	AnnounceResult(winners []string, correctGuess int)
}

var collectors Collectors
var currentCollector Collector

func Init(configPath string, config config.UserCollect, r fiber.Router) {

	r.Get("/", get)
	r.Get("/cancel", cancel)

	twitchChat, name := Chat.New(configPath, r.Group("/chat"))

	collectors = append(collectors, cols{
		Name:      name,
		Collector: twitchChat,
	})

	i := slices.IndexFunc(collectors, func(c cols) bool {
		return c.Name == config.Collector
	})

	if i == -1 {
		return
	}
	currentCollector = collectors[i]

}

func get(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]any{
		"current": currentCollector,
		"all":     collectors,
	})
}

// cancel ends the running collection of guesses
func cancel(ctx *fiber.Ctx) error {
	EndCollect()

	ctx.Status(http.StatusNoContent)
	return nil
}
