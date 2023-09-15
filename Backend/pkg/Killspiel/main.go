package Killspiel

import (
	"Killspiel/pkg/Leaderboard"
	"Killspiel/pkg/ResultCollector"
	"Killspiel/pkg/User"
	"Killspiel/pkg/UserCollector"
	"Killspiel/pkg/config"
	"Killspiel/pkg/database"
	"Killspiel/pkg/router"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"math"
	"net/http"
	"time"
)

const float64EqualityThreshold = 1e-3

var guesses = map[int]UserCollector.Guess{}

func Init(app *fiber.App) {
	path, err := config.FindConfOrDefault()
	if err != nil {
		log.Fatal(err)
	}

	conf, err := config.GetConfig(path)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Init(path)
	if err != nil {
		log.Fatal(err)
	}

	api := router.CreateApiGroup(app)
	UserCollector.Init(path, conf, api.Group("/collector"))
	ResultCollector.Init(path, api.Group("/data"))

	Leaderboard.Init(api.Group("/leaderboard"))
	User.Init(api.Group("/user"))

	// Redirecting user access per link, because /user/:id is routed through svelte
	app.Get("/user/:id/", func(ctx *fiber.Ctx) error {
		ctx.Status(http.StatusFound).Location(fmt.Sprintf("/Redirect/?%s", ctx.OriginalURL()))
		return nil
	})
}

func Run() {
	for Ready() {
		gameInfo := ResultCollector.Begin()

		clear(guesses)

		UserCollector.Collect(guesses)

		var err error
		gameId, err := saveGuesses(gameInfo)
		if err != nil {
			log.Printf("Error saving guesses: %v", err)
			return
		}

		res := ResultCollector.Result()

		winners := getWinners(res, gameId)

		UserCollector.AnnounceResult(res, winners)
	}
}

func Ready() bool {
	for !UserCollector.Ready() {
		time.Sleep(5 * time.Second)
	}
	return true
}

func getWinners(correctGuess float64, gameId int64) []string {
	var winners []string

	tx, err := database.DB.Begin()
	if err != nil {
		log.Printf("Error starting db transaction: %v\n", err)
		return nil
	}
	defer tx.Rollback()

	_, err = database.SetGameCorrect.Exec(correctGuess, gameId)
	if err != nil {
		log.Printf("Error starting db transaction: %v\n", err)
		return nil
	}

	for id, user := range guesses {
		add := 0
		if math.Abs(user.Guess-correctGuess) < float64EqualityThreshold {
			add = 1
			winners = append(winners, user.Name)
		}
		_, err = tx.Stmt(database.UpdateUser).Exec(add, add, id)
		if err != nil {
			log.Printf("Error starting db transaction: %v\n", err)
			continue
		}
	}
	_ = tx.Commit()
	return winners
}

func saveGuesses(dbinfo string) (gameId int64, err error) {
	tx, err := database.DB.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	result, err := tx.Stmt(database.CreateGame).Exec(dbinfo)
	if err != nil {
		return
	}
	if gameId, err = result.LastInsertId(); err != nil {
		err = tx.QueryRow("SELECT last_insert_rowid()").Scan(&gameId)
		if err != nil {
			return
		}
	}

	for id, guess := range guesses {
		res, err := tx.Stmt(database.UpdateUserGuesses).Exec(id)
		if err != nil {
			log.Printf("Error occurd getting player %d: %v\n", id, err)
			continue
		}
		if n, err := res.RowsAffected(); n == 0 && err == nil {
			log.Println(err)
			_, err = tx.Stmt(database.CreateUser).Exec(id, guess.Name)
			if err != nil {
				log.Printf("Error occurd creating player %d: %v\n", id, err)
				continue
			}
		} else if err != nil {
			log.Printf("Error occurd checking if player exists player %d: %v\n", id, err)
			continue
		}

		_, err = tx.Stmt(database.CreateVote).Exec(gameId, id, guess.Guess)
		if err != nil {
			log.Printf("Error occurd getting player %d: %v\n", id, err)
			continue
		}

	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error occured commiting game %d: %v\n", gameId, err)
		return
	}

	return
}
