package users

import (
	"gobackend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(route fiber.Router) {
	route.Get("/users", GetUsers)
	route.Post("/users", middleware.IsSuperAdmin, NewUser)
	route.Delete("/users/:id", middleware.IsSuperAdmin, DeleteUser)
}
