package api

import (
	"gobackend/api/auth"
	"gobackend/api/books"
	"gobackend/api/items"
	"gobackend/api/tasks"
	"gobackend/api/users"
	"gobackend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/api/v1", middleware.Protected())
	authg := app.Group("api/auth")
	books.Routes(v1)
	items.Routes(v1)
	users.Routes(v1)
	tasks.Routes(v1)
	auth.Routes(authg)
}
