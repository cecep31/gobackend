package presenter

import (
	"gobackend/database"
	"gobackend/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type post struct {
	database.DefaultModel
	Title     string    `json:"title"`
	Body      string    `json:"desc"`
	CreatedBy uuid.UUID `json:"created_by"`
}

func PostSuccessResponse(data *entities.Posts) *fiber.Map {
	newData := post{
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

func PostsSuccessResponse(data *[]entities.Posts) *fiber.Map {

	if len(*data) == 0 {
		return &fiber.Map{
			"success": true,
			"data":    []post{},
			"error":   nil,
		}
	}
	return &fiber.Map{
		"success": true,
		"data":    data,
		"error":   nil,
	}
}

func PostErrorResponse(err interface{}) *fiber.Map {
	return &fiber.Map{
		"success": false,
		"data":    "",
		"error":   err,
	}
}
