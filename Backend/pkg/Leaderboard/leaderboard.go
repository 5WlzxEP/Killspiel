package Leaderboard

import (
	"Killspiel/pkg/database"
	"arena"
	"database/sql"
	"github.com/gofiber/fiber/v2"
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
		rows, err = database.LeaderboardDesc.Query(sort, limit, start)
	} else {
		rows, err = database.LeaderboardAsc.Query(sort, limit, start)
	}

	if err != nil {
		return err
	}
	mem := arena.NewArena()
	defer mem.Free()

	// Ranks doesn't start at 0 but at 1
	start++

	var result []User
	result = arena.MakeSlice[User](mem, limit, limit)

	i := 0
	for ; rows.Next() && i < limit; i++ {
		err = rows.Scan(&result[i].Id, &result[i].Name, &result[i].Guesses, &result[i].Points, &result[i].Latest)
		if err != nil {
			rows.Close()
			return err
		}
		result[i].Rank = start + i
	}

	userCount, err := database.GetUserCount()
	if err != nil {
		return err
	}

	return ctx.JSON(map[string]any{
		"data": result[0:i],
		"meta": meta{userCount},
	})
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
