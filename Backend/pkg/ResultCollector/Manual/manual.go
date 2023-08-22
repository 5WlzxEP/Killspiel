package Manual

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type state int

const (
	none state = iota
	begin
	running
	result
)

var WrongState = errors.New("wrong State")

type Manuel struct {
	begin  chan struct{}
	result chan int
	state
}

func New(r fiber.Router) *Manuel {
	m := &Manuel{
		begin:  make(chan struct{}),
		result: make(chan int),
		state:  none,
	}

	r.Get("/", m.get)
	r.Post("/", m.post)

	return m
}

func (m *Manuel) Begin() error {
	if m.state != none {
		return WrongState
	}
	m.state = begin

	<-m.begin

	m.state = running

	return nil
}

func (m *Manuel) Result() (res int, err error) {
	if m.state != running {
		return 0, WrongState
	}
	m.state = result

	res = <-m.result

	m.state = none
	return res, nil
}

func (m *Manuel) get(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]state{"state": m.state})
}

func (m *Manuel) post(ctx *fiber.Ctx) error {
	var data postRequest
	err := ctx.BodyParser(&data)
	if err != nil {
		return err
	}

	switch {
	case m.state == begin && data.Type == "begin":
		m.begin <- struct{}{}

		return ctx.SendStatus(http.StatusNoContent)

	case m.state == result && data.Type == "result":
		m.result <- data.Result

		return ctx.SendStatus(http.StatusNoContent)
	}

	return ctx.SendStatus(http.StatusBadRequest)
}

type postRequest struct {
	Type   string `json:"type"`
	Result int    `json:"result"`
}
