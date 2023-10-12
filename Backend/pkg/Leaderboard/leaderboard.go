package Leaderboard

import (
	"Killspiel/pkg/database"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"sync"
)

type User struct {
	Id      int    `json:"id"`
	Rank    int    `json:"rank"`
	Name    string `json:"name"`
	Guesses int    `json:"guesses"`
	Points  int    `json:"points"`
	Latest  uint8  `json:"latest"`
}

type meta struct {
	Size int `json:"size"`
}

type leaderboard struct {
	meta `json:"meta"`
	Data []User `json:"data"`
}

var leaderboardBuffer = sync.Pool{New: func() any { return leaderboard{Data: make([]User, 100)} }}

func get(ctx *fiber.Ctx) error {

	limit := ctx.QueryInt("l", 25)
	start := ctx.QueryInt("p", 0)

	sort := parseSort(ctx.Query("s", "p"))
	order := ctx.QueryInt("o", 0)

	limit = min(limit, 100)

	var (
		err  error
		rows *sql.Rows
	)
	if order == 0 {
		rows, err = database.LeaderboardDesc.QueryContext(ctx.Context(), sort, limit, start)
	} else {
		rows, err = database.LeaderboardAsc.QueryContext(ctx.Context(), sort, limit, start)
	}
	if err != nil {
		return err
	}

	board := leaderboardBuffer.Get().(leaderboard)
	defer leaderboardBuffer.Put(board)

	// Ranks don't start at 0 but at 1
	start++

	i := 0
	for ; rows.Next() && i < limit; i++ {
		err = rows.Scan(&board.Data[i].Id, &board.Data[i].Name, &board.Data[i].Guesses, &board.Data[i].Points, &board.Data[i].Latest)
		if err != nil {
			_ = rows.Close()
			return err
		}
		board.Data[i].Rank = start + i
	}
	_ = rows.Close()

	board.Data = board.Data[:i]

	board.Size, err = database.GetUserCount(ctx.Context())
	if err != nil {
		return err
	}

	return ctx.JSON(board)
}

func Init(router fiber.Router) {
	router.Get("/", get)
}

func parseSort(s string) int {
	switch s {
	case "n":
		return 2
	case "g":
		return 3
	case "p":
		return 1
	}
	return 1
}
