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
	"sync"
	"time"
)

type key struct {
	Key string `json:"key,omitempty"`
	id  int
}

var buf = helper.Buffer[key]{}

// DO NOT CHANGE BUFFER SIZE, or only accordingly to database.GetUserVotes
var historyPool = sync.Pool{New: func() any { return User{History: make([]Game, 25)} }}

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
	Id        int       `json:"id"`
	Guess     float64   `json:"guess"`
	Correct   float64   `json:"correct"`
	Precision float64   `json:"precision"`
	Time      time.Time `json:"time"`
}

func get(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	user := historyPool.Get().(User)
	defer historyPool.Put(user)
	user.Id = id

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

	i := 0
	for ; rows.Next(); i++ {
		err = rows.Scan(&user.History[i].Id, &user.History[i].Correct, &user.History[i].Guess, &user.History[i].Time, &user.History[i].Precision)
		if err != nil {
			i--
			continue
		}
	}

	user.History = user.History[:i]

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
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Points int    `json:"points"`
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
		err = rows.Scan(&res.Id, &res.Name, &res.Points)
		if err != nil {
			continue
		}
		result = append(result, res)
	}

	return ctx.JSON(result)
}
