package ws

import (
	"gobackend/middleware"
	"gobackend/pkg/entities"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

var connections []*websocket.Conn

func WsSetup(app *fiber.App) {
	app.Get("/ws/global", middleware.Protectedws(), middleware.GetUser, websocket.New(func(c *websocket.Conn) {
		user := c.Locals("datauser").(entities.Users)

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

		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				log.Println("WebSocket read error:", err)
				break
			}

			h, errjson := json.Marshal(&fiber.Map{"msg": string(message), "userID": user.ID})
			if errjson != nil {
				return
			}
			log.Printf("WebSocket message: %s", []byte(h))
			// Broadcast the incoming message to all other WebSocket connections
			for _, conn := range connections {
				if err := conn.WriteMessage(messageType, []byte(h)); err != nil {
					log.Println("WebSocket write error:", err)
					break
				}
			}
		}
	}))
}
