package items

import (
	"errors"

	"github.com/cecep31/gobackend/database"
	"github.com/cecep31/gobackend/pkg"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

func GetItem(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var item Items
	err := db.First(&item, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkg.EntityNotFound("Item not found")
	} else if err != nil {
		return pkg.Unexpected(err.Error())
	}
	return c.JSON(item)
}
