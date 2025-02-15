package presenter

import (
	"gobackend/database/template"
	"gobackend/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type post struct {
	template.DefaultModel
	Title     string    `json:"title"`
	Body      string    `json:"desc"`
	CreatedBy uuid.UUID `json:"created_by"`
}

func PostSuccess(postData *entities.Post) *fiber.Map {
	resp := &fiber.Map{
		"success": true,
		"data": fiber.Map{
			"title":      postData.Title,
			"desc":       postData.Body,
			"created_by": postData.CreatedBy,
		},
		"error": nil,
	}

	return resp
}
func PostsSuccessResponse(data *[]entities.Post) *fiber.Map {

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
