package presenter

import (
	"gobackend/pkg/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Image        string    `json:"image"`
	Issuperadmin bool      `json:"issuperadmin" gorm:"default:false"`
	CreateAt     time.Time `json:"createAt"`
}

func UserSuccessResponse(data *entities.User) interface{} {
	user := User{
		ID:           data.ID,
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		Email:        data.Email,
		Image:        data.Image,
		Issuperadmin: data.Issuperadmin,
		CreateAt:     data.CreatedAt,
	}
	return user
}

func UsersSuccessResponse(data *[]entities.User) interface{} {
	if len(*data) == 0 {
		return &fiber.Map{
			"success": true,
			"data":    []User{},
			"error":   nil,
		}
	}
	var newData []User
	for _, item := range *data {
		newUser := User{
			ID:           item.ID,
			FirstName:    item.FirstName,
			LastName:     item.LastName,
			Email:        item.Email,
			Image:        item.Image,
			Issuperadmin: item.Issuperadmin,
		}
		newData = append(newData, newUser)
	}

	return newData
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"success": false,
		"data":    "",
		"error":   err.Error(),
	}
}
