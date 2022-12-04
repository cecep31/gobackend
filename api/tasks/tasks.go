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
