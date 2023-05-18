package handlers

import (
	"gobackend/api/presenter"
	"gobackend/pkg/entities"
	"gobackend/pkg/user"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Users
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		result, err := service.InserUser(&requestBody)
		if err != nil {
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(result))
	}
}
func GetUsers(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, err := service.GetUsers()
		if err != nil {
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UsersSuccessResponse(user))
	}
}
func GetUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idparam := c.Params("id")
		id, err := uuid.Parse(idparam)
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		user, err := service.GetUser(id)
		if err != nil {
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UserSuccessResponse(user))
	}
}
