package handlers

import (
	"gobackend/api/presenter"
	"gobackend/pkg/entities"
	"gobackend/pkg/posts"
	"gobackend/pkg/utils"
	"log"
	"os"
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

		resulvalidate := utils.ValidateThis(requestBody)
		if resulvalidate != nil {
			return c.Status(422).JSON(presenter.ErrorResponse(resulvalidate, "data is not valite"))
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

func UpdatePost(service posts.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		param_post_id := c.Params("id")
		var requestBody posts.Posts
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.PostErrorResponse(err.Error()))
		}
		post_id, _ := uuid.Parse(param_post_id)
		requestBody.ID = post_id

		resulvalidate := utils.ValidateThis(requestBody)
		if resulvalidate != nil {
			return c.JSON(presenter.ErrorResponse(resulvalidate))
		}

		if err != nil {
			return c.JSON(presenter.PostErrorResponse(err.Error()))
		}
		errrepo := service.UpdatePost(&requestBody)
		if errrepo != nil {
			return c.JSON(presenter.PostErrorResponse(err))
		}
		return c.Status(fiber.StatusCreated).JSON(&requestBody)
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
				return c.Status(fiber.StatusBadRequest).SendString("Only one file allowed")
			}
			file := files[0]
			allowedExtensions := []string{".jpg", ".jpeg", ".png"}
			if !service.ValidFileExtension(file.Filename, allowedExtensions) {
				return c.Status(fiber.StatusBadRequest).SendString("Invalid file extension")
			}
			filename, errput := service.PutObjectPhoto(c.Context(), "post_photo/"+file.Filename, file)
			if errput != nil {
				return c.Status(500).JSON(presenter.ErrorResponse(errput.Error()))
			}
			return c.JSON(fiber.Map{
				"photo_url": os.Getenv("STORAGE_URL") + "/" + filename,
			})
		}
		return c.Status(500).JSON(presenter.ErrorResponse("not form"))
	}
}
