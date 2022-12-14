package api

import (
	"gobackend/api/auth"
	"gobackend/api/books"
	"gobackend/api/items"
	"gobackend/api/posts"
	"gobackend/api/tasks"
	"gobackend/api/users"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	v1 := api.Group("/v1")
	// v1noauth := app.Group("/api/v1")
	authg := app.Group("api/auth")
	books.Routes(v1)
	items.Routes(v1)
	users.Routes(v1)
	tasks.Routes(v1)
	auth.Routes(authg)
	posts.Routes(v1)
	// payments.Routes(v1)

}
