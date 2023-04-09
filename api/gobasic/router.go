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
	route.Post("/testbodyparse", func(c *fiber.Ctx) error {
		var data interface{}
		if err := c.BodyParser(&data); err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"hello": data,
		})
	})
}
