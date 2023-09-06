package UserCollector

import (
	"Killspiel/pkg/UserCollector/Chat"
	"Killspiel/pkg/config"
	"context"
	"github.com/gofiber/fiber/v2"
	"slices"
	"time"
)

type Collectors []cols

type cols struct {
	Name string `json:"name"`
	Collector
}

type Collector interface {
	Ready() bool
	CollectGuesses(ctx context.Context, collector func(id int, user, guess string))
	AnnounceResult(winners []string, correctGuess float64)
}

var collectors Collectors
var currentCollector Collector
var collectTime time.Duration

func Init(configPath string, config config.UserCollect, r fiber.Router) {

	r.Get("/", get)

	twitchChat, name := Chat.New(configPath, r.Group("/chat"))

	collectors = append(collectors, cols{
		Name:      name,
		Collector: twitchChat,
	})

	collectTime = config.Duration

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
		"time":    collectTime,
	})
}
