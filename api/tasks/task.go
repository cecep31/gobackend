package tasks

import (
	"gobackend/database"
	"gobackend/pkg"
	"gobackend/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	db := database.DB
	var taks []entities.Tasks
	db.Preload("Group").Find(&taks)
	return c.JSON(taks)
}

func GetTaskGroups(c *fiber.Ctx) error {
	db := database.DB
	var taskgroups []entities.Taskgorups
	db.Preload("Task").Find(&taskgroups)
	return c.JSON(taskgroups)
}

func GetMyTaskGroup(c *fiber.Ctx) error {
	user := c.Locals("datauser").(entities.Users)
	db := database.DB
	var taskgroups []entities.Taskgorups
	db.Where("created_by = ?", user.ID).Preload("Task").Find(&taskgroups)
	return c.JSON(taskgroups)
}
func GetTaskGroup(c *fiber.Ctx) error {
	db := database.DB
	var taskgroup entities.Taskgorups
	db.Preload("Task").First(&taskgroup)
	return c.JSON(taskgroup)
}

func GetMyTasks(c *fiber.Ctx) error {
	user := c.Locals("datauser").(entities.Users)
	var task []entities.Tasks
	println(user.Email)
	db := database.DB
	db.Where("created_by = ?", user.ID).Find(&task)
	return c.JSON(task)
}

func NewTaskGroup(c *fiber.Ctx) error {
	type Taskgorups struct {
		database.DefaultModel
		Name  string `json:"name"`
		Order int64  `json:"order"`
	}
	user := c.Locals("datauser").(entities.Users)
	// _ := c.Locals("datauser").(entities.User)
	db := database.DB
	taskgroup := new(Taskgorups)
	if err := c.BodyParser(taskgroup); err != nil {
		return pkg.BadRequest("Invalid params")
	}
	taskgroupnew := new(entities.Taskgorups)
	taskgroupnew.Name = taskgroup.Name
	taskgroupnew.Order = taskgroup.Order
	taskgroupnew.Created_by = user.ID

	// taskgroup.Created_by = int64(user.ID)
	db.Create(&taskgroupnew)
	return c.JSON(taskgroupnew)
}

func NewTask(c *fiber.Ctx) error {
	user := c.Locals("datauser").(entities.Users)
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
	var tasknew entities.Tasks
	tasknew.Title = task.Title
	tasknew.Desc = task.Desc
	tasknew.GroupID = task.GroupID
	tasknew.Created_by = user.ID

	err2 := db.Create(&tasknew).Error
	if err2 != nil {
		return pkg.BadRequest("Ivalid Request")
	}
	return c.Status(200).JSON(tasknew)
}
