package Manual

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type state int

const (
	none state = iota
	result
	begin
)

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

func (m *Manuel) Begin() {
	m.state = begin

	<-m.begin

	m.state = none
}

func (m *Manuel) Result() (res int) {
	m.state = result

	res = <-m.result

	m.state = none
	return res
}

func (m *Manuel) get(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]state{"State": m.state})
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
