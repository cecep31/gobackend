package api

import (
	"github.com/cecep31/gobackend/api/auth"
	"github.com/cecep31/gobackend/api/books"
	"github.com/cecep31/gobackend/api/items"
	"github.com/cecep31/gobackend/api/users"
	"github.com/cecep31/gobackend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/api/v1", middleware.Protected())
	authg := app.Group("api/auth")
	books.Routes(v1)
	items.Routes(v1)
	users.Routes(v1)
	auth.Routes(authg)
}
