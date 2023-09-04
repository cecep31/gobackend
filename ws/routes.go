package ws

import (
	"gobackend/pkg/posts"
	"gobackend/ws/handlers"

	"github.com/gofiber/fiber/v2"
)

func WsPostRouter(app fiber.Router, service posts.Service) {
	app.Get("/posts/:slug", handlers.Comments(service))
}
