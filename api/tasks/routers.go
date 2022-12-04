package tasks

import "github.com/gofiber/fiber/v2"

func Routes(route fiber.Router) {
	route.Get("tasks", GetTasks)
	route.Get("mytasks", GetMyTasks)
}
