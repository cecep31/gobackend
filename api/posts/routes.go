package posts

import (
	"gobackend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(route fiber.Router) {
	route.Post("/posts", middleware.Protected(), Newpost)
	route.Get("/posts", GetPosts)
}
