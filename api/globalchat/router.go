package globalchat

import "github.com/gofiber/fiber/v2"

func Routes(route fiber.Router) {
	route.Get("/globalchat", Globalchat)
}
