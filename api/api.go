package api

import (
	"github.com/cecep31/gobackend/api/books"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/api/v1")
	books.Routes(v1)
}
