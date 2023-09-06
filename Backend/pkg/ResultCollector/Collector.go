package ResultCollector

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type Collector interface {
	// Begin blocks until the Condition for the Collector to realize the voting event should start
	Begin(ctx context.Context, cancelFunc context.CancelFunc)
	// Result blocks until a result is present and returns it
	Result(ctx context.Context, c chan float64)
}

type States uint8

type state struct {
	States `json:"state"`
}

var (
	State           state
	beginCancelFunc context.CancelFunc
	resultChan      = make(chan float64)
	collector       Collector
)

const (
	none States = iota
	begin
	running
	result
)

func Init(r fiber.Router) {
	r.Get("/", get)
	r.Post("/", post)
}

func get(ctx *fiber.Ctx) error {
	return ctx.JSON(State)
}

func Begin() {
	var ctx context.Context
	ctx, beginCancelFunc = context.WithCancel(context.Background())
	State.States = none
	if collector != nil {
		collector.Begin(ctx, beginCancelFunc)
	}

	<-ctx.Done()
	State.States = begin
}

func Result() float64 {
	for len(resultChan) > 0 {
		<-resultChan
	}

	State.States = running

	ctx, cancel := context.WithCancel(context.Background())
	if collector != nil {
		collector.Result(ctx, resultChan)
	}
	res := <-resultChan
	cancel()

	State.States = result
	return res
}
