package presenter

import (
	"gobackend/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Book is the presenter object which will be passed in the response by Handler
type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	Image        string    `json:"image"`
	Issuperadmin bool      `json:"issuperadmin" gorm:"default:false"`
}

// BookSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func UserSuccessResponse(data *entities.Users) *fiber.Map {
	book := User{
		ID:           data.ID,
		Email:        data.Email,
		Username:     data.Username,
		Image:        data.Image,
		Issuperadmin: data.Issuperadmin,
	}
	return &fiber.Map{
		"status": true,
		"data":   book,
		"error":  nil,
	}
}

// BooksSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func UsersSuccessResponse(data *[]entities.Users) *fiber.Map {
	var newData []User
	for _, item := range *data {
		newUser := User{
			ID:           item.ID,
			Email:        item.Email,
			Username:     item.Username,
			Image:        item.Image,
			Issuperadmin: item.Issuperadmin,
		}
		newData = append(newData, newUser)
	}
	return &fiber.Map{
		"status": true,
		"data":   newData,
		"error":  nil,
	}
}

// BookErrorResponse is the ErrorResponse that will be passed in the response by Handler
func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
