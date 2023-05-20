package api

import (
	"gobackend/api/handlers"
	"gobackend/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service user.Service) {
	app.Post("/users", handlers.AddUser(service))
	app.Get("/users", handlers.GetUsers(service))
	app.Get("/users/:id", handlers.GetUser(service))
}

func AuthRouter(app fiber.Router) {
	app.Get("/oauth", handlers.Loginoatuth)
	app.Get("/oauth/callback", handlers.CallbackHandler)
}
