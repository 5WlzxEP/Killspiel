package UserCollector

import (
	"Killspiel/pkg/database"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var EndCollect context.CancelFunc
var guesses map[int]database.Guess

func Ready() bool {
	if currentCollector != nil {
		return currentCollector.Collector.Ready()
	}
	return false
}

func Collect() {

	clear(guesses)

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
		guesses[id] = database.Guess{
			Name:  user,
			Guess: g,
		}
	})

	var err error
	_, err = database.SaveGuesses(&guesses)
	if err != nil {
		log.Printf("Error saving guesses: %v", err)
		return
	}

}

func AnnounceResult(correctGuess float64) {
	var winners []string

	tx, err := database.DB.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()
	for id, user := range guesses {
		add := 0
		if user.Guess == correctGuess {
			add = 1
			winners = append(winners, user.Name)
		}
		_, err = tx.Stmt(database.UpdateUser).Exec(add, add, id)
		if err != nil {
			return
		}
	}
	_ = tx.Commit()

	if currentCollector != nil {
		currentCollector.AnnounceResult(winners, correctGuess)
	} else {
		_, _ = fmt.Fprintln(os.Stderr, "No UserCollector defined, skipping announcing winners")
	}
}
