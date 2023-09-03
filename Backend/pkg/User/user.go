package User

import (
	"Killspiel/pkg/database"
	"Killspiel/pkg/helper"
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type key struct {
	Key string `json:"key,omitempty"`
	id  int
}

var buf = helper.Buffer[key]{}

func Init(r fiber.Router) {
	r.Get("/:id/", get)
	r.Delete("/:id/", del)
	r.Post("/", search)
}

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Points  int    `json:"points"`
	Guesses int    `json:"guesses"`
	Latest  int    `json:"latest"`
	History []Game `json:"history"`
}

type Game struct {
	Id      int     `json:"id"`
	Guess   float64 `json:"guess"`
	Correct float64 `json:"correct"`
}

func get(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	user := User{Id: id, History: make([]Game, 0, 50)}

	err = database.GetUser.QueryRow(id).Scan(&user.Name, &user.Points, &user.Guesses, &user.Latest)
	if errors.Is(err, sql.ErrNoRows) {
		return ctx.SendStatus(404)
	}
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
		buf.Add(k)
		return ctx.JSON(k)
	}

	if buf.Contains(k) {
		go deleteUser(id)
		return ctx.SendStatus(http.StatusNoContent)
	}

	return ctx.SendStatus(http.StatusBadRequest)
}

func getKey() string {
	return uuid.NewString()
}

func deleteUser(id int) {
	_, err := database.DeleteUser.Exec(id)
	if err != nil {
		log.Printf("An error occured deleteing user %d: %v\n", id, err)
	}
}

type searchName struct {
	Name string `json:"name"`
}

type searchResult struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func search(ctx *fiber.Ctx) error {
	var name searchName
	_ = ctx.BodyParser(&name)
	if len(name.Name) < 3 {
		return ctx.JSON(fiber.Map{
			"error": "name must be at least 3 characters long",
		})
	}

	var result []searchResult

	rows, err := database.SearchUser.Query("%" + name.Name + "%")
	if err != nil {
		return err
	}

	for rows.Next() {
		var res searchResult
		err = rows.Scan(&res.Id, &res.Name)
		if err != nil {
			continue
		}
		result = append(result, res)
	}

	return ctx.JSON(result)
}
