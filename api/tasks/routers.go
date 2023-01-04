package tasks

import (
	"gobackend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(route fiber.Router) {
	route.Get("tasks", middleware.Protected(), middleware.IsSuperAdmin, GetTasks)
	route.Get("mytasks", middleware.Protected(), middleware.GetUser, GetMyTasks)
	route.Post("taskgroups", middleware.Protected(), middleware.GetUser, NewTaskGroup)
	route.Get("taskgroups", middleware.Protected(), middleware.IsSuperAdmin, GetTaskGroups)
	route.Post("tasks", middleware.Protected(), middleware.GetUser, NewTask)
	route.Get("mytaskgroups", middleware.Protected(), middleware.GetUser, GetMyTaskGroup)
}
