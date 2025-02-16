package handlers

import (
	"gobackend/api/presenter"
	"gobackend/pkg/entities"
	"gobackend/pkg/posts"
	"gobackend/pkg/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func AddPost(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody posts.PostCreate
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.PostErrorResponse(err.Error()))
		}

		resulvalidate := utils.ValidateThis(&requestBody)
		if len(resulvalidate) > 0 {
			return c.Status(422).JSON(presenter.ErrorResponse(resulvalidate, "data is not valite"))
		}

		userlocal := c.Locals("user").(*jwt.Token)
		claims := userlocal.Claims.(jwt.MapClaims)
		useridstring := claims["id"].(string)
		userid, err := uuid.Parse(useridstring)

		if err != nil {
			return c.JSON(presenter.PostErrorResponse(err.Error()))
		}

		realpost := entities.Post{Title: requestBody.Title, Body: requestBody.Body, CreatedBy: userid, Slug: requestBody.Slug, Photo_url: requestBody.Photo_url}
		result, err := service.InserPosts(&realpost)
		if err != nil {
			return c.JSON(presenter.PostErrorResponse(err))
		}
		return c.Status(fiber.StatusCreated).JSON(result)
	}
}

func UpdatePost(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		param_post_id := c.Params("id")
		var requestBody posts.PostUpdate
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.PostErrorResponse(err.Error()))
		}
		post_id, _ := uuid.Parse(param_post_id)
		requestBody.ID = post_id

		resulvalidate := utils.ValidateThis(requestBody)
		if len(resulvalidate) > 0 {
			return c.JSON(presenter.ErrorResponse(resulvalidate))
		}
		errrepo := service.UpdatePost(&requestBody)
		if errrepo != nil {
			return c.JSON(presenter.PostErrorResponse(err))
		}
		return c.Status(fiber.StatusCreated).JSON(&requestBody)
	}
}

func GetPostsRandom(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		postsdata, err := service.GetPostsRandom()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   err.Error(),
				"success": false,
				"data":    nil,
				"message": "error",
			})
		}
		return c.JSON(fiber.Map{
			"data":    postsdata,
			"error":   nil,
			"success": true,
			"message": "success",
		})
	}
}

func GetPosts(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		offset := c.Query("offset", "0")
		limit := c.Query("limit", "100")

		// Convert query parameters to integers
		offsetInt, _ := strconv.Atoi(offset)
		limitInt, _ := strconv.Atoi(limit)

		totalPosts, err := service.GetTotalPosts()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		postsdata, err := service.GetPostsPaginated(offsetInt, limitInt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		response := fiber.Map{
			"total": totalPosts,
			"data":  postsdata,
		}
		return c.JSON(response)
	}
}
func GetPost(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		slug := c.Params("slug")
		post, err := service.GetPost(slug)
		if err != nil {
			return c.JSON(presenter.PostErrorResponse(err))
		}
		return c.JSON(post)
	}
}

func DeletePost(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		_, err := service.GetPostByid(id)
		if err != nil {
			return c.Status(404).JSON(presenter.ErrorResponse("Post not Found"))
		}
		errd := service.DeletePost(id)
		if errd != nil {
			return c.Status(500).JSON(presenter.ErrorResponse(errd.Error()))
		}
		return c.SendStatus(200)
	}
}

func UploadPhotoHandler(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if form, err := c.MultipartForm(); err == nil {
			// => *multipart.Form

			// Get all files from "documents" key:
			files := form.File["image"]
			if len(files) != 1 {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(fiber.Map{"image": "Only one file allowed"}))
			}
			file := files[0]
			allowedExtensions := []string{".jpg", ".jpeg", ".png"}
			if !service.ValidFileExtension(file.Filename, allowedExtensions) {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(fiber.Map{"image": "Invalid file extention"}))
			}
			if file.Size > 2<<20 {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(fiber.Map{"image": "File size exceeds the limit (2MB)"}))
			}

			userlocal := c.Locals("user").(*jwt.Token)
			claims := userlocal.Claims.(jwt.MapClaims)
			useridstring := claims["id"].(string)
			userid, _ := uuid.Parse(useridstring)

			filename, errput := service.UploadPhoto(c.Context(), file, userid)
			if errput != nil {
				return c.Status(500).JSON(presenter.ErrorResponse(errput.Error()))
			}
			return c.JSON(fiber.Map{
				"photo_url": filename,
			})
		}
		return c.Status(500).JSON(presenter.ErrorResponse("not form"))
	}
}
