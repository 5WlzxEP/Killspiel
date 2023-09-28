package Game

import (
	"Killspiel/pkg/database"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Game struct {
	Id           int       `json:"id"`
	Correct      float64   `json:"correct"`
	Time         time.Time `json:"time"`
	CorrectCount int       `json:"correctCount"`
	UserCount    int       `json:"userCount"`
	Precision    float64   `json:"precision"`
	Verteilung   string    `json:"verteilung"`
}

func Init(r fiber.Router) {
	r.Get("/:id/", get)
	r.Get("/:id/:vote/", getVotes)
}

func get(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	g := Game{Id: id}
	err = database.GetGame.QueryRow(id).Scan(&g.Correct, &g.Time, &g.CorrectCount, &g.UserCount, &g.Precision, &g.Verteilung)
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

	rows, err := database.GetPlayersByVote.Query(id, vote, vote)
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
