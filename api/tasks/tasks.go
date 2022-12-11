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
	db.Preload("Group").Find(&taks)
	return c.JSON(taks)
}

func GetTaskGroups(c *fiber.Ctx) error {
	db := database.DB
	var taskgroup []entities.Taskgorup
	db.Find(&taskgroup)
	return c.JSON(taskgroup)
}

func GetMyTaskGroup(c *fiber.Ctx) error {
	user := c.Locals("datauser").(entities.User)
	db := database.DB
	var taskgroup []entities.Taskgorup
	db.Where("created_by = ?", user.ID).Preload("Task").Find(&taskgroup)
	return c.JSON(taskgroup)

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

func NewTask(c *fiber.Ctx) error {
	user := c.Locals("datauser").(entities.User)
	db := database.DB
	type Task struct {
		database.DefaultModel
		Title   string `json:"title"`
		Desc    string `json:"desc" gorm:"type:text"`
		GroupID int64  `json:"group_id"`
		Order   int64  `json:"order"`
	}
	task := new(Task)
	if err := c.BodyParser(task); err != nil {
		return pkg.BadRequest("Invalid params")
		// return c.SendString(err)
	}
	var tasknew entities.Task
	tasknew.Title = task.Title
	tasknew.Desc = task.Desc
	tasknew.GroupID = task.GroupID
	tasknew.Created_by = int64(user.ID)

	db.Create(&tasknew)
	return c.JSON(tasknew)
}
