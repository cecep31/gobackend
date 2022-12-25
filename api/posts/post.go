package posts

import (
	"gobackend/database"
	"gobackend/pkg"
	"gobackend/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

func Newpost(c *fiber.Ctx) error {
	db := database.DB
	post := new(entities.Posts)
	if err := c.BodyParser(post); err != nil {
		return pkg.BadRequest("invalid params")
	}

	result := db.Create(&post).Error
	// return result
	if result != nil {
		return pkg.BadRequest("invalid params")
	}
	return c.JSON(post)
}

func GetPosts(c *fiber.Ctx) error {
	db := database.DB
	var posts []entities.Posts
	db.Find(&posts)
	return c.JSON(posts)

}
