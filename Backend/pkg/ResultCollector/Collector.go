package ResultCollector

import (
	"Killspiel/pkg/ResultCollector/RiotApi"
	"Killspiel/pkg/Websocket"
	"Killspiel/pkg/config"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"math"
	"time"
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

	// Prefix returns the prefix of the info dbInfo sends in Begin
	Prefix() string
	// Resolve is called during startup if Prefix matches with the dbinfo string from Begin
	Resolve(dbinfo string) (float64, error)
}

type States uint8

func setState(state States) {
	State = state
	Websocket.Broadcast([]byte{'S', 't', 'a', 't', 'e', ':', ' ', '0' + uint8(state)})
}

var (
	State            States
	beginCancelFunc  context.CancelFunc
	resultChan       = make(chan float64)
	currentCollector Collector
	collectors       []*collector
	dbInfo           = make(chan string, 5)
	conf             *config.Config
	endTime          int64
)

const (
	none States = iota
	begin
	running
	result
)

func Init(confPath string, conf2 *config.Config, r fiber.Router) {
	r.Get("/", get)
	r.Post("/", post)

	c, n := RiotApi.New(confPath, r.Group("/riot"))
	collectors = append(collectors, &collector{
		Collector: c,
		Name:      n,
	})

	currentCollector = c
	conf = conf2
}

func get(ctx *fiber.Ctx) error {
	if State == begin {
		return ctx.JSON(struct {
			State States `json:"state"`
			End   int64  `json:"end"` // end of collection votes
		}{State, endTime})
	}
	return ctx.JSON(struct {
		State States `json:"state"`
	}{State})
}

func Begin(ctx context.Context) string {
	// clear dbInfo
	for len(dbInfo) > 0 {
		<-dbInfo
	}

	var c context.Context
	c, beginCancelFunc = context.WithCancel(ctx)

	setState(none)

	if currentCollector != nil && currentCollector.Ready() {
		go currentCollector.Begin(c, beginCancelFunc, dbInfo)
	}

	// still blocks if no currentCollector exists, which is ok, because the user can manually override it
	<-c.Done()

	select {
	case <-ctx.Done():
		return ""
	default:
	}

	//setState(begin)
	State = begin
	endTime = time.Now().Add(conf.UserCollector.Duration * time.Second).Unix()
	Websocket.Broadcast([]byte(fmt.Sprintf("State: 1, %v", endTime)))

	select {
	case x := <-dbInfo:
		return x
	default:
		return "manual"
	}
}

func Result(ctx context.Context) float64 {
	for len(resultChan) > 0 {
		<-resultChan
	}

	setState(running)

	ctx, cancel := context.WithCancel(ctx)
	if currentCollector != nil && currentCollector.Ready() {
		go currentCollector.Result(ctx, resultChan)
	}
	var res float64
	select {
	case res = <-resultChan:
	case <-ctx.Done():
		res = math.NaN()
	}
	cancel()

	setState(result)

	return res
}
