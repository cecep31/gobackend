package posts

import (
	"gobackend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(route fiber.Router) {
	route.Post("/posts", middleware.Protected(), middleware.GetUser, Newpost)
	route.Get("/posts", GetPosts)
	route.Get("/posts/:id", getpost)
}
