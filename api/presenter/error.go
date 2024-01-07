package presenter

import "github.com/gofiber/fiber/v2"

func ErrorResponse(error interface{}, message ...string) *fiber.Map {
	var msg string
	if len(message) > 0 {
		msg = message[0]
	} else {
		msg = "_"
	}
	return &fiber.Map{
		"success": false,
		"message": msg,
		"error":   error,
	}
}
