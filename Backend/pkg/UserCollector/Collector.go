package UserCollector

import (
	"Killspiel/pkg/UserCollector/Chat"
	"context"
	"github.com/gofiber/fiber/v2"
)

var collectors []Collector

type Collector interface {
	Ready() bool
	CollectGuesses(ctx context.Context) map[string]int
	AnnounceResult(winners []string, correctGuess int)
}

func Init(configPath string, r fiber.Router) {
	twitchChat := Chat.New(configPath, r.Group("/chat"))

	collectors = append(collectors, twitchChat)
}
