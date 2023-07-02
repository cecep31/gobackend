package api

import (
	"gobackend/api/handlers"
	"gobackend/middleware"
	"gobackend/pkg/auth"
	"gobackend/pkg/post"
	"gobackend/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service user.Service) {
	app.Post("/users", handlers.AddUser(service))
	app.Get("/users", handlers.GetUsers(service))
	app.Get("/users/:id", handlers.GetUser(service))
}

func AuthRouter(app fiber.Router, service auth.Service) {
	app.Get("/oauth", handlers.Loginoatuth)
	app.Get("/oauth/callback", handlers.CallbackHandler(service))
}

func PostRouter(app fiber.Router, service post.Service) {
	app.Get("/posts", handlers.GetPosts(service))
	app.Get("/posts/:id", handlers.GetPost(service))
	app.Post("/posts", middleware.Protected(), handlers.AddPost(service))
}
