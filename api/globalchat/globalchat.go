package globalchat

import (
	"gobackend/database"
	"gobackend/pkg"
	"gobackend/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

func Globalchat(c *fiber.Ctx) error {
	db := database.DB
	var data []entities.Globalchat
	result := db.Preload("User").Find(&data)
	if result.Error != nil {
		return pkg.Unexpected(result.Error.Error())
	}
	return c.JSON(data)
}
