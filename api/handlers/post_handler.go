package handlers

import (
	"gobackend/api/presenter"
	"gobackend/pkg/entities"
	"gobackend/pkg/posts"
	"gobackend/pkg/validator"
	"log"
	"strconv"

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
			return c.JSON(presenter.ErrorResponse(resulvalidate))
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
		random := c.Query("random")
		if random == "true" {
			posts, err := service.GetPostsRandom()
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			return c.JSON(posts)
		}
		page := c.Query("page", "1")
		itemsPerPage := c.Query("per_page", "5")

		// Convert query parameters to integers
		pageInt, _ := strconv.Atoi(page)
		perPageInt, _ := strconv.Atoi(itemsPerPage)

		totalPosts, err := service.GetTotalPosts()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		posts, err := service.GetPostsPaginated(pageInt, perPageInt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		response := fiber.Map{
			"total": totalPosts,
			"page":  pageInt,
			"data":  posts,
		}
		return c.JSON(response)
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
