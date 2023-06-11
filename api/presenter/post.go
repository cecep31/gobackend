package presenter

import (
	"github.com/google/uuid"
	"gobackend/database"
	"gobackend/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

type Posts struct {
	database.DefaultModel
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	CreatedBy uuid.UUID `json:"created_by"`
}

func PostSuccessResponse(data *entities.Posts) *fiber.Map {
	newData := Posts{
		Title:     data.Title,
		Desc:      data.Desc,
		CreatedBy: data.CreatedBy,
	}
	return &fiber.Map{
		"status": true,
		"data":   newData,
		"error":  nil,
	}
}

func PostsSuccessResponse(data *[]entities.Posts) *fiber.Map {

	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func PostErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
