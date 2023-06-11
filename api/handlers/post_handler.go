package handlers

import (
	"gobackend/api/presenter"
	"gobackend/pkg/entities"
	"gobackend/pkg/post"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddPost(service post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Posts
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		result, err := service.InserPosts(&requestBody)
		if err != nil {
			return c.JSON(presenter.PostErrorResponse(err))
		}
		return c.JSON(presenter.PostSuccessResponse(result))
	}
}
func GetPosts(service post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		post, err := service.GetPosts()
		if err != nil {
			return c.JSON(presenter.PostErrorResponse(err))
		}
		return c.JSON(presenter.PostsSuccessResponse(post))
	}
}
func GetPost(service post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idparam := c.Params("id")
		id, err := uuid.Parse(idparam)
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		post, err := service.GetPost(id)
		if err != nil {
			return c.JSON(presenter.PostErrorResponse(err))
		}
		return c.JSON(presenter.PostSuccessResponse(post))
	}
}
