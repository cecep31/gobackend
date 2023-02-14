package ws

import "github.com/gofiber/fiber/v2"

func WsSetup(app *fiber.App) {
	app.Group("/ws")
}
