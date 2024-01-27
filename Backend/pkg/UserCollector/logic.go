package UserCollector

import (
	"Killspiel/pkg/Websocket"
	"bytes"
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
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

func Collect(ctx context.Context, guesses map[int]Guess) {

	deadline := time.Now().Add(collectTime)
	ctx, EndCollect = context.WithDeadline(ctx, deadline)

	if currentCollector == nil {
		_, _ = fmt.Fprintln(os.Stderr, "No UserCollector defined, skipping collecting guesses")
		return
	}

	currentCollector.CollectGuesses(ctx, func(id int, user, guess string, color ...string) {
		g, err := strconv.ParseFloat(guess, 64)
		if err != nil {
			return
		}
		guesses[id] = Guess{
			Name:  user,
			Guess: g,
		}

		c := ""
		if len(color) > 0 {
			c = color[0]
		}

		data, err := json.Marshal(fiber.Map{"name": user, "vote": guess, "color": c})
		if err != nil {
			return
		}
		Websocket.Broadcast(bytes.Join([][]byte{[]byte("Vote: "), data}, nil))
	})
}

func AnnounceResult(correctGuess float64, winners []string) {

	if currentCollector != nil {
		currentCollector.AnnounceResult(winners, correctGuess)
	} else {
		_, _ = fmt.Fprintln(os.Stderr, "No UserCollector defined, skipping announcing winners")
	}
}
