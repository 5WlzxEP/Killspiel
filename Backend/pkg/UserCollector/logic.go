package UserCollector

import (
	"Killspiel/pkg/config"
	"Killspiel/pkg/database"
	"context"
	"log"
	"strconv"
	"time"
)

type Guess struct {
	Name  string
	Guess float64
}

var EndCollect context.CancelFunc
var guesses map[int]Guess

func Collect(config config.UserCollect) {

	clear(guesses)

	deadline := time.Now().Add(config.Duration * time.Second)
	var ctx context.Context
	ctx, EndCollect = context.WithDeadline(context.Background(), deadline)

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

	err := database.SaveGuesses(&guesses)
	if err != nil {
		log.Printf("Error saving guesses: %v", err)
		return
	}

}
