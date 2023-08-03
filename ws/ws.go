package ws

import (
	"gobackend/database"
	"gobackend/middleware"
	"gobackend/pkg/entities"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var connections []*websocket.Conn

func WsSetup(app *fiber.App) {
	app.Get("/ws/global", middleware.Protectedws(), middleware.GetUser, websocket.New(func(c *websocket.Conn) {
		user := c.Locals("datauser").(entities.Users)
		db := database.DB

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

			var globalchat = new(entities.Globalchat)
			globalchat.ID = uuid.New()
			globalchat.UserID = user.ID
			globalchat.Msg = string(message)

			dberr := db.Create(&globalchat)
			if dberr.Error != nil {
				log.Printf("error fb: %s", dberr.Error)
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
