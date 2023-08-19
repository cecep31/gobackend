package handlers

import (
	"gobackend/api/presenter"
	"gobackend/pkg/entities"
	"gobackend/pkg/posts"
	"gobackend/pkg/validator"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func AddPost(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody posts.Posts
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.PostErrorResponse(err.Error()))
		}

		resulvalidate := validator.ValidateThis(requestBody)
		if resulvalidate != nil {
			return c.JSON(presenter.PostErrorResponse(resulvalidate))
		}

		userlocal := c.Locals("user").(*jwt.Token)
		claims := userlocal.Claims.(jwt.MapClaims)
		useridstring := claims["id"].(string)
		userid, err := uuid.Parse(useridstring)

		if err != nil {
			return c.JSON(presenter.PostErrorResponse(err.Error()))
		}

		realpost := entities.Posts{Title: requestBody.Title, Body: requestBody.Body, CreatedBy: userid}
		result, err := service.InserPosts(&realpost)
		if err != nil {
			return c.JSON(presenter.PostErrorResponse(err))
		}
		return c.Status(fiber.StatusCreated).JSON(result)
	}
}
func GetPosts(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		posts, err := service.GetPosts()
		if err != nil {
			return c.JSON(presenter.PostErrorResponse(err))
		}
		return c.JSON(posts)
	}
}
func GetPost(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		slug := c.Params("slug")
		log.Println(slug)
		post, err := service.GetPost(slug)
		if err != nil {
			return c.JSON(presenter.PostErrorResponse(err))
		}
		return c.JSON(post)
	}
}
