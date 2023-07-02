package api

import (
	"gobackend/api/auth"
	"gobackend/api/globalchat"
	"gobackend/api/gobasic"
	"gobackend/api/tasks"
	"gobackend/api/users"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	v1 := api.Group("/v1")
	authg := app.Group("auth")
	users.Routes(v1)
	tasks.Routes(v1)
	auth.Routes(authg)
	gobasic.Routes(v1)
	globalchat.Routes(v1)
	// payments.Routes(v1)

}
