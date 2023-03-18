package ws

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var connections []*websocket.Conn

func WsSetup(app *fiber.App) {
	app.Get("/ws/global", websocket.New(func(c *websocket.Conn) {
		connections = append(connections, c)
		defer func() {
			log.Println("WebSocket disconnected")
			for i, conn := range connections {
				if conn == c {
					connections = append(connections[:i], connections[i+1:]...)
					break
				}
			}
		}()
		fmt.Println(connections)

		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				log.Println("WebSocket read error:", err)
				break
			}
			log.Printf("WebSocket message: %s", message)

			// Broadcast the incoming message to all other WebSocket connections
			for _, conn := range connections {
				if conn != c {
					if err := conn.WriteMessage(messageType, message); err != nil {
						log.Println("WebSocket write error:", err)
						break
					}
				}
			}
		}
	}))
}
