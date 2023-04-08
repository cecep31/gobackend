package gobasic

import "github.com/gofiber/fiber/v2"

func Routes(route fiber.Router) {
	route.Get("/testpointer", func(c *fiber.Ctx) error {
		data := 200
		newdata := data
		return c.JSON(fiber.Map{
			"old data": data,
			"new data": newdata,
		})
	})
}
