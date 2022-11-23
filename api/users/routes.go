package users

import "github.com/gofiber/fiber/v2"

func Routes(route fiber.Router) {
	route.Get("/users", GetUsers)
	route.Post("/users", NewUser)
	route.Delete("/users/:id", DeleteUser)
}
