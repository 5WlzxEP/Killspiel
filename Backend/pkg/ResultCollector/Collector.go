package ResultCollector

import (
	"Killspiel/pkg/ResultCollector/RiotApi"
	"Killspiel/pkg/Websocket"
	"context"
	"github.com/gofiber/fiber/v2"
)

type collector struct {
	Collector `json:"-"`
	Name      string `json:"name"`
}

type Collector interface {
	Ready() bool
	// Begin blocks until the Condition for the Collector to realize the voting event should start
	Begin(ctx context.Context, cancelFunc context.CancelFunc, dbInfo chan<- string)
	// Result blocks until a result is present and returns it
	Result(ctx context.Context, c chan float64)
}

type States uint8

type state struct {
	States `json:"state"`
}

var (
	State            state
	beginCancelFunc  context.CancelFunc
	resultChan       = make(chan float64)
	currentCollector Collector
	collectors       []*collector
	dbInfo           = make(chan string, 5)
)

const (
	none States = iota
	begin
	running
	result
)

func Init(conf string, r fiber.Router) {
	r.Get("/", get)
	r.Post("/", post)

	c, n := RiotApi.New(conf, r.Group("/riot"))
	collectors = append(collectors, &collector{
		Collector: c,
		Name:      n,
	})

	currentCollector = c
}

func get(ctx *fiber.Ctx) error {
	return ctx.JSON(State)
}

func Begin() string {
	// clear dbInfo
	for len(dbInfo) > 0 {
		<-dbInfo
	}

	var ctx context.Context
	ctx, beginCancelFunc = context.WithCancel(context.Background())

	setState(none)

	if currentCollector != nil && currentCollector.Ready() {
		go currentCollector.Begin(ctx, beginCancelFunc, dbInfo)
	}

	// still blocks if no currentCollector exists, which is ok, because the user can manually override it
	<-ctx.Done()

	setState(begin)

	select {
	case x := <-dbInfo:
		return x
	default:
		return "manual"
	}
}

func Result() float64 {
	for len(resultChan) > 0 {
		<-resultChan
	}

	setState(running)

	ctx, cancel := context.WithCancel(context.Background())
	if currentCollector != nil && currentCollector.Ready() {
		go currentCollector.Result(ctx, resultChan)
	}
	res := <-resultChan
	cancel()

	setState(result)

	return res
}
