package tasks

import (
	"gobackend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(route fiber.Router) {
	route.Get("tasks", GetTasks)
	route.Get("mytasks", middleware.GetUser, GetMyTasks)
	route.Post("taskgroups", middleware.GetUser, NewTaskGroup)
	route.Post("tasks", middleware.GetUser, NewTask)
	route.Get("mytaskgroups", middleware.GetUser, GetMyTaskGroup)
}
