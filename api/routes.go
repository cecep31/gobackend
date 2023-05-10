package api

import (
	"gobackend/api/handlers"
	"gobackend/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service user.Service) {
	app.Post("/users", handlers.AddUser(service))
}
