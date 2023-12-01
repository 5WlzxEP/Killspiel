package Websocket

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"time"
)

var ws = map[*websocket.Conn]chan struct{}{}
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
		ws[c] = ch
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
	for conn, ch := range ws {
		err := conn.WriteMessage(msgType, msg)
		if err != nil {
			delete(ws, conn)
			close(ch)
		}
	}
}
