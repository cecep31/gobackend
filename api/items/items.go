package items

import (
	"github.com/cecep31/gobackend/database"
	"github.com/gofiber/fiber/v2"
)

type Items struct {
	database.DefaultModel
	Name  string `json:"name"`
	Desc  string `json:"desc" gorm:"type:text"`
	Image string `json:"image" gorm:"type:text"`
}

func GetItems(c *fiber.Ctx) error {
	db := database.DB
	var items []Items
	db.Find(&items)
	println(len(items))
	for i := 0; i < len(items); i++ {
		println("wk")
	}

	return c.JSON(items)
}
