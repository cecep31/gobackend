package presenter

import "github.com/gofiber/fiber/v2"

func ErrorResponse(error interface{}) *fiber.Map {
	return &fiber.Map{
		"success": false,
		"error":   error,
	}
}
