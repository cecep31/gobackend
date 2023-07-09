package presenter

import (
	"gobackend/database"
	"gobackend/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Posts struct {
	database.DefaultModel
	Title     string    `json:"title"`
	Body      string    `json:"desc"`
	CreatedBy uuid.UUID `json:"created_by"`
}

func TaskSuccessResponse(data *entities.Tasks) *fiber.Map {
	newData := Posts{
		Title:     data.Title,
		Body:      data.Body,
		CreatedBy: data.CreatedBy,
	}
	return &fiber.Map{
		"success": true,
		"data":    newData,
		"error":   nil,
	}
}

func TasksSuccessResponse(data *[]entities.Tasks) *fiber.Map {

	if len(*data) == 0 {
		return &fiber.Map{
			"success": true,
			"data":    []Posts{},
			"error":   nil,
		}
	}
	return &fiber.Map{
		"success": true,
		"data":    data,
		"error":   nil,
	}
}

func TaskErrorResponse(err interface{}) *fiber.Map {
	return &fiber.Map{
		"success": false,
		"data":    "",
		"error":   err,
	}
}
