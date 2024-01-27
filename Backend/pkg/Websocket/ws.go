package Websocket

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"sync"
	"time"
)

var (
	ws       = map[*websocket.Conn]chan struct{}{}
	lock     = sync.Mutex{}
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
		lock.Lock()
		ws[c] = ch
		lock.Unlock()
		<-ch
		channels.Put(ch)
	}))

	go broadcast()

}

func Close() {
	lock.Lock()
	defer lock.Unlock()
	for conn, c := range ws {
		delete(ws, conn)
		c <- struct{}{}
	}
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
	lock.Lock()
	defer lock.Unlock()
	for conn, ch := range ws {
		err := conn.WriteMessage(msgType, msg)
		if err != nil {
			delete(ws, conn)
			ch <- struct{}{}
		}
	}
}
