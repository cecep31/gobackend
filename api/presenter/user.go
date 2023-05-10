package presenter

import (
	"gobackend/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Book is the presenter object which will be passed in the response by Handler
type Book struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}

// BookSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func BookSuccessResponse(data *entities.Users) *fiber.Map {
	book := Book{
		ID:       data.ID,
		Email:    data.Email,
		Username: data.Username,
	}
	return &fiber.Map{
		"status": true,
		"data":   book,
		"error":  nil,
	}
}

// BooksSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func BooksSuccessResponse(data *[]Book) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// BookErrorResponse is the ErrorResponse that will be passed in the response by Handler
func BookErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
