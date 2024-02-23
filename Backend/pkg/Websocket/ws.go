package Websocket

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"sync"
	"time"
)

var (
	ws       = sync.Map{}
	sendCh   = make(chan []byte)
	channels = sync.Pool{New: func() any { return make(chan struct{}, 1) }}
)

func Init(r fiber.Router) {
	r.Use("/", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	r.Get("/", websocket.New(func(c *websocket.Conn) {

		defer c.Close()
		//ch := make(chan struct{})
		ch := channels.Get().(chan struct{})
		ws.Store(c, ch)
		<-ch
		channels.Put(ch)
	}))

	go broadcast()

}

func Close() {
	ws.Range(func(key, value any) bool {
		ws.Delete(key)
		value.(chan struct{}) <- struct{}{}
		return true
	})
}

func Broadcast(msg []byte) {
	sendCh <- msg
}

func broadcast() {
	ticker := time.NewTicker(1 * time.Minute)

	for {
		select {
		case msg := <-sendCh:
			send(1, msg)
		case <-ticker.C:
			send(9, nil)
		}
	}
}

func send(msgType int, msg []byte) {
	ws.Range(func(key, value any) bool {
		err := key.(*websocket.Conn).WriteMessage(msgType, msg)
		if err != nil {
			ws.Delete(key)
			value.(chan struct{}) <- struct{}{}
		}
		return true
	})
}
