package posts

import (
	"errors"
	"gobackend/database"
	"gobackend/pkg"
	"gobackend/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Newpost(c *fiber.Ctx) error {
	user := c.Locals("datauser").(entities.Users)
	type postvalidate struct {
		Title string `json:"title"`
		Desc  string `json:"desc"`
	}
	validate := new(postvalidate)
	db := database.DB
	post := new(entities.Posts)
	if err := c.BodyParser(validate); err != nil {
		return pkg.BadRequest("invalid params")
	}
	post.Title = validate.Title
	post.Desc = validate.Desc
	post.Created_by = int64(user.ID)

	result := db.Create(&post).Error
	// return result
	if result != nil {
		return pkg.BadRequest("invalid params db")
	}
	return c.JSON(post)
}

func GetPosts(c *fiber.Ctx) error {
	db := database.DB
	var posts []entities.Posts
	db.Find(&posts)
	return c.JSON(posts)
}

func getpost(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var post entities.Posts
	err := db.First(&post, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkg.EntityNotFound("record Not Found")
	} else if err != nil {
		return pkg.Unexpected(err.Error())
	}
	return c.JSON(post)
}
