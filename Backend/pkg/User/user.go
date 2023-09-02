package User

import (
	"Killspiel/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func Init(r fiber.Router) {
	r.Get("/:id/", get)
	r.Delete("/:id/", del)
}

type User struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Points  int    `json:"points,omitempty"`
	Guesses int    `json:"guesses,omitempty"`
	Latest  int    `json:"latest,omitempty"`
	History []Game `json:"history,omitempty"`
}

type Game struct {
	Id      int     `json:"id,omitempty"`
	Guess   float64 `json:"guess,omitempty"`
	Correct float64 `json:"correct,omitempty"`
}

func get(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	user := User{Id: id, History: make([]Game, 0, 50)}

	err = database.GetUser.QueryRow(id).Scan(user.Name, user.Points, user.Guesses, user.Latest)
	if err != nil {
		return err
	}

	rows, err := database.GetUserGames.Query(id)
	if err != nil {
		return err
	}
	for rows.Next() {
		g := Game{}
		err = rows.Scan(g.Id, g.Correct, g.Guess)
		if err != nil {
			continue
		}

		user.History = append(user.History, g)
	}
	return ctx.JSON(user)
}

// del requires to steps to delete to avoid accidentally deletes
func del(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	k := key{id: id}
	err = ctx.BodyParser(&k)
	if err != nil || k.Key == "" {
		k.Key = getKey()
		Buffer.add(k)
		return ctx.JSON(k)
	}

	if Buffer.contains(k) {
		go func() {
			_, err := database.DeleteUser.Exec(id)
			if err != nil {
				log.Printf("An error occured deleteing user %d: %v\n", id, err)
			}
		}()
		return ctx.SendStatus(http.StatusNoContent)
	}

	return ctx.SendStatus(http.StatusBadRequest)
}

func getKey() string {
	return uuid.NewString()
}
