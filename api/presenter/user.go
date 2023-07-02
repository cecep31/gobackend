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
	Username     string    `json:"username"`
	Image        string    `json:"image"`
	Issuperadmin bool      `json:"issuperadmin" gorm:"default:false"`
	CreateAt     time.Time `json:"createAt"`
}

func UserSuccessResponse(data *entities.Users) *fiber.Map {
	user := User{
		ID:           data.ID,
		Email:        data.Email,
		Image:        data.Image,
		Issuperadmin: data.Issuperadmin,
		CreateAt:     data.CreatedAt,
	}
	return &fiber.Map{
		"status": true,
		"data":   user,
		"error":  nil,
	}
}

func UsersSuccessResponse(data *[]entities.Users) *fiber.Map {
	var newData []User
	for _, item := range *data {
		newUser := User{
			ID:           item.ID,
			Email:        item.Email,
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

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
