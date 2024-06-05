package Killspiel

import (
	"Killspiel/pkg/Game"
	"Killspiel/pkg/Leaderboard"
	"Killspiel/pkg/ResultCollector"
	"Killspiel/pkg/User"
	"Killspiel/pkg/UserCollector"
	"Killspiel/pkg/Websocket"
	"Killspiel/pkg/config"
	"Killspiel/pkg/database"
	"Killspiel/pkg/router"
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"math"
	"net/http"
	"time"
)

var guesses = map[int]UserCollector.Guess{}
var verteilung = map[int]int{}
var conf *config.Config

// Init initializes the Killspiel
func Init(app *fiber.App) {
	path, err := config.FindConfOrDefault()
	if err != nil {
		logger.Fatal(err)
	}

	conf, err = config.GetConfig(path)
	if err != nil {
		logger.Fatal(err)
	}

	err = database.Init(path)
	if err != nil {
		logger.Fatal(err)
	}

	api := router.CreateApiGroup(app)

	api.Get("/", get)
	api.Post("/", post)

	UserCollector.Init(path, conf, api.Group("/collector"))
	ResultCollector.Init(path, conf, api.Group("/data"))

	Leaderboard.Init(api.Group("/leaderboard"))
	User.Init(api.Group("/user"))
	Game.Init(api.Group("/game"), conf)

	games := ResultCollector.CheckUnfinishedGames()
	for _, v := range games {
		getWinners(v.Correct, v.GameId)
	}

	// Redirecting user access per link, because /user/:id is routed through svelte
	app.Get("/user/:id/", func(ctx *fiber.Ctx) error {
		ctx.Status(http.StatusMovedPermanently).Location(fmt.Sprintf("/Redirect/?%s", ctx.OriginalURL()))
		return nil
	})
	app.Get("/game/", func(ctx *fiber.Ctx) error {
		ctx.Status(http.StatusMovedPermanently).Location(fmt.Sprintf("/Redirect/?%s", ctx.OriginalURL()))
		return nil
	})
}

// Run starts the Killspiel
func Run(ctx context.Context, finished chan<- struct{}) {

	canceled := getCancel(ctx)

	for Ready(canceled) && !canceled() {
		gameInfo := ResultCollector.Begin(ctx)

		// shutdown
		if gameInfo == "" {
			break
		}

		clear(guesses)

		UserCollector.Collect(ctx, guesses)

		var err error
		gameId, err := saveGuesses(gameInfo)
		if err != nil {
			logger.Printf("Error saving guesses: %v", err)
			continue
		}

		res := ResultCollector.Result(ctx)

		// shutdown
		if math.IsNaN(res) {
			break
		}

		winners := getWinners(res, gameId)

		UserCollector.AnnounceResult(res, winners)
	}

	// shutdown of the websockets
	Websocket.Close()
	close(finished)
}

// getCancel returns a function that returns true if the context is canceled
func getCancel(ctx context.Context) func() bool {
	done := ctx.Done()

	return func() bool {
		select {
		case <-done:
			return true
		default:
			return false
		}
	}
}

// Ready waits until the UserCollector is ready
func Ready(canceled func() bool) bool {
	for !canceled() && !UserCollector.Ready() {
		time.Sleep(5 * time.Second)
	}
	return true
}

// getWinners returns the winner names and saves the result to the database
func getWinners(correctGuess float64, gameId int64) []string {
	var winners []string

	tx, err := database.DB.Begin()
	if err != nil {
		logger.Printf("Error starting db transaction: %v\n", err)
		return nil
	}
	defer tx.Rollback()

	for id, user := range guesses {
		add := 0
		if math.Abs(user.Guess-correctGuess) <= conf.Precision {
			add = 1
			winners = append(winners, user.Name)
		}
		_, err = tx.Stmt(database.UpdateUser).Exec(add, add, id)
		if err != nil {
			logger.Printf("Error updating User %d: %v\n", id, err)
			continue
		}
	}

	_, err = tx.Stmt(database.SetGameCorrect).Exec(correctGuess, conf.Precision, len(winners), gameId)
	if err != nil {
		logger.Printf("Error updating game %d: %v\n", gameId, err)
	}

	_ = tx.Commit()
	return winners
}

// saveGuesses saves the guesses to the database and returns the gameId
func saveGuesses(dbInfo string) (gameId int64, err error) {
	tx, err := database.DB.Begin()
	if err != nil {
		logger.Printf("Error starting db transaction: %v\n", err)
		return
	}
	defer tx.Rollback()

	result, err := tx.Stmt(database.CreateGame).Exec(len(guesses), dbInfo)
	if err != nil {
		return
	}
	if gameId, err = result.LastInsertId(); err != nil {
		err = tx.QueryRow("SELECT last_insert_rowid()").Scan(&gameId)
		if err != nil {
			return
		}
	}

	clear(verteilung)

	for id, guess := range guesses {
		verteilung[int(math.Floor(guess.Guess))]++
		res, err := tx.Stmt(database.UpdateUserGuesses).Exec(id)
		if err != nil {
			logger.Printf("Error occurred getting player %d: %v\n", id, err)
			continue
		}
		if n, err := res.RowsAffected(); n == 0 && err == nil {
			_, err = tx.Stmt(database.CreateUser).Exec(id, guess.Name)
			if err != nil {
				logger.Printf("Error occurred creating player %d: %v\n", id, err)
				continue
			}
		} else if err != nil {
			logger.Printf("Error occurred checking if player exists player %d: %v\n", id, err)
			continue
		}

		_, err = tx.Stmt(database.CreateVote).Exec(gameId, id, guess.Guess)
		if err != nil {
			logger.Printf("Error occurred getting player %d: %v\n", id, err)
			continue
		}

	}

	bytes, err := json.Marshal(verteilung)
	if err != nil {
		logger.Printf("An error occurred marshaling verteilung: %v\n", err)
	} else {
		_, err = tx.Stmt(database.SetGameVerteilung).Exec(string(bytes), gameId)
		if err != nil {
			logger.Printf("An error occured saving Verteilung to DB: %v\n", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		logger.Printf("Error occured commiting game %d: %v\n", gameId, err)
	}

	return
}
