package tasks

import (
	"gobackend/database"
	"gobackend/pkg"
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

func NewTaskGroup(c *fiber.Ctx) error {
	type Taskgorup struct {
		database.DefaultModel
		Name  string `json:"name"`
		Order int64  `json:"order"`
	}
	user := c.Locals("datauser").(entities.User)
	// _ := c.Locals("datauser").(entities.User)
	db := database.DB
	taskgroup := new(Taskgorup)
	if err := c.BodyParser(taskgroup); err != nil {
		return pkg.BadRequest("Invalid params")
	}
	taskgroupnew := new(entities.Taskgorup)
	taskgroupnew.Name = taskgroup.Name
	taskgroupnew.Order = taskgroup.Order
	taskgroupnew.Created_by = int64(user.ID)

	// taskgroup.Created_by = int64(user.ID)
	db.Create(&taskgroupnew)
	return c.JSON(taskgroupnew)
}
