package tasks

import (
	"gobackend/database"
	"gobackend/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	db := database.DB
	var taks []entities.Task
	db.Find(&taks)
	return c.JSON(taks)
}
func GetMyTasks(c *fiber.Ctx) error {
	user := c.Locals("datauser").(entities.User)
	var task []entities.Task
	println(user.Email)
	db := database.DB
	db.Where("created_by = ?", user.ID).Find(&task)
	return c.JSON(task)
}
