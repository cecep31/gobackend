package auth

import "github.com/gofiber/fiber/v2"

func Routes(route fiber.Router) {
	route.Post("/login", Login)
}
