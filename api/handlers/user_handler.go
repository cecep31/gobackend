package handlers

import (
	"gobackend/api/presenter"
	"gobackend/pkg/entities"
	"gobackend/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func AddUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Users
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.BookErrorResponse(err))
		}
		result, err := service.InserUser(&requestBody)
		if err != nil {
			return c.JSON(presenter.BookErrorResponse(err))
		}
		return c.JSON(presenter.BookSuccessResponse(result))
	}
}
