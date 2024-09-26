package handlers

import (
	"gobackend/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func GetWriter(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.GetWriter()

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusCreated).JSON(result)
	}
}
