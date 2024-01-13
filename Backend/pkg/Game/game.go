package Game

import (
	"Killspiel/pkg/config"
	"Killspiel/pkg/database"
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var pool = sync.Pool{New: func() any { return make([]Game, 50) }}
var conf *config.Config

type Game struct {
	Id           int       `json:"id"`
	Correct      *float64  `json:"correct"`
	Time         time.Time `json:"time"`
	CorrectCount int       `json:"correctCount"`
	UserCount    int       `json:"userCount"`
	Precision    float64   `json:"precision"`
	Verteilung   string    `json:"verteilung,omitempty"`
}

func Init(r fiber.Router, c *config.Config) {
	r.Get("/", get)
	r.Get("/latest/", latest)
	r.Get("/:id/", getId)
	r.Get("/:id/:vote/", getVotes)
	r.Put("/:id/", setResult)

	conf = c
}

func get(ctx *fiber.Ctx) error {
	limit := min(ctx.QueryInt("limit", 50), 50)
	offset := ctx.QueryInt("offset", 0)

	res := pool.Get().([]Game)

	rows, err := database.GetLatestGames.QueryContext(ctx.Context(), limit, offset)
	if err != nil {
		return err
	}
	defer rows.Close()

	i := 0
	for ; rows.Next() && i < 50; i++ {
		err = rows.Scan(&res[i].Id, &res[i].Correct, &res[i].UserCount, &res[i].CorrectCount, &res[i].Precision, &res[i].Time)
		if err != nil {
			log.Println(err)
			i--
			continue
		}
	}

	return ctx.JSON(res[:i])
}

func getId(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	g := Game{Id: id}
	err = database.GetGame.QueryRowContext(ctx.Context(), id).Scan(&g.Correct, &g.Time, &g.CorrectCount, &g.UserCount, &g.Precision, &g.Verteilung)
	if err != nil {
		return err
	}

	return ctx.JSON(g)
}

type User struct {
	Name string  `json:"name"`
	Id   int     `json:"id"`
	Vote float64 `json:"vote"`
}

func getVotes(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	vote, err := ctx.ParamsInt("vote")
	if err != nil {
		return err
	}

	var result []User

	rows, err := database.GetPlayersByVote.QueryContext(ctx.Context(), id, vote, vote)
	if err != nil {
		return err
	}

	for rows.Next() {
		var res User
		err = rows.Scan(&res.Id, &res.Name, &res.Vote)
		if err != nil {
			continue
		}
		result = append(result, res)
	}

	return ctx.JSON(result)
}

func latest(ctx *fiber.Ctx) error {
	var g Game

	err := database.GetLastGame.QueryRowContext(ctx.Context()).Scan(&g.Id, &g.Correct, &g.UserCount, &g.CorrectCount, &g.Precision, &g.Time, &g.Verteilung)
	if err != nil {
		return err
	}

	return ctx.JSON(g)
}

func setResult(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	res, err := strconv.ParseFloat(string(ctx.Body()), 64)
	if err != nil {
		return err
	}

	err = database.CheckGameUnfinished.QueryRow(id).Scan(&id)
	if err != nil {
		return ctx.Status(http.StatusExpectationFailed).SendString("Result is already set")
	}

	err = getWinners(res, id)
	if err != nil {
		return err
	}

	return ctx.SendStatus(http.StatusNoContent)
}

func getWinners(correctGuess float64, gameId int) error {

	tx, err := database.DB.Begin()
	if err != nil {
		log.Printf("Error starting db transaction: %v\n", err)
		return nil
	}
	defer tx.Rollback()

	rows, err := tx.Stmt(database.GetUsersByGame).Query(gameId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			_, err = tx.Stmt(database.SetGameCorrect).Exec(correctGuess, conf.Precision, 0, gameId)
			if err != nil {
				log.Printf("Error updating game %d: %v\n", gameId, err)
				return err
			}

			return nil
		}
		return err
	}
	defer rows.Close()

	winners := 0

	var guess float64
	var userId int64
	for rows.Next() {
		err = rows.Scan(&userId, &guess)
		if err != nil {
			return err
		}

		add := 0
		if math.Abs(guess-correctGuess) <= conf.Precision {
			add = 1
			winners++
		}
		_, err = tx.Stmt(database.UpdateUser).Exec(add, add, userId)
		if err != nil {
			log.Printf("Error updating User %d: %v\n", userId, err)
			continue
		}
	}

	_, err = tx.Stmt(database.SetGameCorrect).Exec(correctGuess, conf.Precision, winners, gameId)
	if err != nil {
		log.Printf("Error updating game %d: %v\n", gameId, err)
		return err
	}

	_ = tx.Commit()
	return nil
}
