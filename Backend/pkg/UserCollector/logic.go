package UserCollector

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Guess struct {
	Name  string
	Guess float64
}

var EndCollect context.CancelFunc

func Ready() bool {
	if currentCollector != nil {
		return currentCollector.Collector.Ready()
	}
	return false
}

func Collect(guesses map[int]Guess) {

	deadline := time.Now().Add(collectTime)
	var ctx context.Context
	ctx, EndCollect = context.WithDeadline(context.Background(), deadline)

	if currentCollector == nil {
		_, _ = fmt.Fprintln(os.Stderr, "No UserCollector defined, skipping collecting guesses")
		return
	}

	currentCollector.CollectGuesses(ctx, func(id int, user, guess string) {
		g, err := strconv.ParseFloat(guess, 64)
		if err != nil {
			return
		}
		guesses[id] = Guess{
			Name:  user,
			Guess: g,
		}
	})
}

func AnnounceResult(correctGuess float64, winners []string) {

	if currentCollector != nil {
		currentCollector.AnnounceResult(winners, correctGuess)
	} else {
		_, _ = fmt.Fprintln(os.Stderr, "No UserCollector defined, skipping announcing winners")
	}
}
