package books

import (
	"gobackend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(route fiber.Router) {
	route.Get("/books", GetBooks)
	route.Get("/books/:id", GetBook)
	route.Put("/books/:id", middleware.Protected(), UpdateBook)
	route.Post("/books", middleware.Protected(), middleware.GetUser, NewBook)
	route.Delete("/books/:id", DeleteBook)
}
