package UserCollector

import "context"

type Collector interface {
	Ready() error
	CollectGuesses(ctx context.Context) map[string]float64
	AnnounceResult(winners []string, correctGuess float32)
}
