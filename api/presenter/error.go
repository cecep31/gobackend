package presenter

import "github.com/gofiber/fiber/v2"

func ErrorResponse(error interface{}, message ...string) *fiber.Map {
	if len(message) == 0 {
		message = append(message, "_") // Default value
	}
	return &fiber.Map{
		"success": false,
		"message": message,
		"error":   error,
	}
}
