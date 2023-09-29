package Websocket

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"time"
)

//type Websockets DoubleLinkedList[*websocket.Conn]

var ws = map[*websocket.Conn]chan struct{}{}

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

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index

		defer c.Close()
		ch := make(chan struct{})
		ws[c] = ch
		<-ch
	}))

	go func() {
		for ; ; time.Sleep(1 * time.Minute) {
			broadcast(9, nil)
		}
	}()
}

func Broadcast(msg []byte) {
	broadcast(1, msg)
}

func broadcast(msgType int, msg []byte) {
	for conn, ch := range ws {
		err := conn.WriteMessage(msgType, msg)
		if err != nil {
			delete(ws, conn)
			close(ch)
		}
	}
}
