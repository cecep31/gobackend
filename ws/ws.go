package ws

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func WsSetup(app *fiber.App) {
	app.Get("/ws/global", websocket.New(func(c *websocket.Conn) {
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			if string(msg) == "hello" {
				if err = c.WriteMessage(mt, []byte("hello juga")); err != nil {
					log.Println("write:", err)
					break
				}
			} else if string(msg) == "hi" {
				if err = c.WriteMessage(mt, []byte("hi juga")); err != nil {
					log.Println("write:", err)
					break
				}
			}

			log.Printf("recv: %s", msg)

			if err = c.WriteMessage(mt, []byte("hello")); err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))
}
