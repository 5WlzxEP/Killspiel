package Websocket

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"sync"
	"time"
)

var ws = map[*websocket.Conn]chan struct{}{}
var lock = sync.RWMutex{}
var sendCh = make(chan []byte)

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
		ch := make(chan struct{})
		lock.Lock()
		ws[c] = ch
		lock.Unlock()
		<-ch
	}))

	go broadcast()

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
	lock.RLock()
	defer lock.RUnlock()
	for conn, ch := range ws {
		err := conn.WriteMessage(msgType, msg)
		if err != nil {
			lock.RUnlock()
			lock.Lock()
			delete(ws, conn)
			close(ch)
			lock.Unlock()
			lock.RLock()
		}
	}
}
