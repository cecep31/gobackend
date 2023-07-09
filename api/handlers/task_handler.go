package handlers

import (
	"gobackend/api/presenter"
	"gobackend/pkg/tasks"

	"github.com/gofiber/fiber/v2"
)

func GetTasks(service tasks.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tasks, err := service.GetTasks()
		if err != nil {
			return c.JSON(presenter.TaskErrorResponse(err))
		}
		return c.JSON(presenter.TasksSuccessResponse(tasks))
	}
}
