package users

import (
	"gobackend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(route fiber.Router) {
	route.Get("/users", middleware.Protected(), GetUsers)
	route.Get("/profile", middleware.Protected(), middleware.GetUser, Getyou)
	route.Post("/users", middleware.Protected(), middleware.IsSuperAdmin, middleware.GetUser, NewUser)
	// route.Post("/users", NewUser)
	route.Delete("/users/:id", middleware.Protected(), middleware.IsSuperAdmin, DeleteUser)
	route.Put("/users/:id", middleware.Protected(), UpdateUser)
	route.Get("users/:id", middleware.Protected(), Getuser)
}
