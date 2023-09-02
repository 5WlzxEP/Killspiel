package UserCollector

import (
	"Killspiel/pkg/config"
	"Killspiel/pkg/database"
	"context"
	"log"
	"strconv"
	"time"
)

var EndCollect context.CancelFunc
var guesses map[int]database.Guess
var latestId int64

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
		guesses[id] = database.Guess{
			Name:  user,
			Guess: g,
		}
	})

	var err error
	latestId, err = database.SaveGuesses(&guesses)
	if err != nil {
		log.Printf("Error saving guesses: %v", err)
		return
	}

}
