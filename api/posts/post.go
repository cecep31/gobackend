package posts

import (
	"gobackend/database"
	"gobackend/pkg"
	"gobackend/pkg/entities"

	"github.com/gofiber/fiber/v2"
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
