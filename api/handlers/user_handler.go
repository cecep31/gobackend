package handlers

import (
	"gobackend/api/presenter"
	"gobackend/pkg/user"
	"gobackend/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody user.Users
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		resultevalidate := utils.ValidateThis(requestBody)
		if len(resultevalidate) > 0 {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(resultevalidate))
		}

		result, err := service.InserUser(&requestBody)
		if err != nil {
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.Status(fiber.StatusCreated).JSON(result)
	}
}
func GetUsers(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, err := service.GetUsers()
		if err != nil {
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(user)
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
		return c.JSON(user)
	}
}

func DeletUser(service user.Service) fiber.Handler {
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
		error := service.DeleteUser(user)
		if error != nil {
			return c.Status(500).JSON(presenter.ErrorResponse(error))
		}
		return c.SendStatus(200)
	}
}
