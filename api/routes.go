package api

import (
	"gobackend/api/handlers"
	"gobackend/middleware"
	"gobackend/pkg/auth"
	"gobackend/pkg/posts"
	"gobackend/pkg/tasks"
	"gobackend/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service user.Service) {
	app.Post("users", handlers.AddUser(service))
	app.Get("users", handlers.GetUsers(service))
	app.Get("users/:id", handlers.GetUser(service))
	app.Put("users/:id", middleware.Protected(), middleware.IsSuperAdmin, func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusServiceUnavailable)
	})
	app.Delete("users/:id", middleware.Protected(), middleware.IsSuperAdmin, handlers.DeletUser(service))

}

func AuthRouter(app fiber.Router, service auth.Service) {
	app.Get("oauth", handlers.Loginoatuth)
	app.Get("oauth/callback", handlers.CallbackHandler(service))
	app.Post("login", handlers.Login)
	app.Get("profile", middleware.Protected(), handlers.Profile(service))
	app.Put("profile", middleware.Protected(), handlers.UpdateProfile(service))
}

func PostRouter(app fiber.Router, service posts.Service) {
	app.Get("posts", handlers.GetPosts(service))
	app.Get("posts/:slug", handlers.GetPost(service))
	app.Post("posts", middleware.Protected(), middleware.IsSuperAdmin, handlers.AddPost(service))
	app.Put("posts/:id", middleware.Protected(), middleware.IsSuperAdmin, handlers.UpdatePost(service))
}

func TaskRouter(app fiber.Router, service tasks.Service) {
	app.Get("tasks", handlers.GetTasks(service))
}
