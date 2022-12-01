package users

import (
	"errors"

	"gobackend/database"
	"gobackend/pkg"
	"gobackend/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// users without paasword
type Users struct {
	database.DefaultModel
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Image    string `json:"image" gorm:"type:text"`
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []Users
	db.Find(&users)
	println(len(users))
	return c.JSON(users)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func NewUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(entities.User)
	if err := c.BodyParser(user); err != nil {
		return pkg.BadRequest("Invalid params")
	}

	var existuser Users
	err := db.Where("username = ?", user.Username).First(&existuser).Error
	if !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return c.JSON(fiber.Map{
			"message": "user telah ada",
		})
	}

	hash, _ := HashPassword(user.Password)
	// println(&bytes)
	// newuser.Password = string(hash)
	user.Password = hash
	db.Create(&user)

	return c.JSON(fiber.Map{
		"username": user.Username,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var user entities.User
	err := db.First(&user, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkg.EntityNotFound("No user found")
	} else if err != nil {
		return pkg.Unexpected(err.Error())
	}
	db.Delete(&user)
	return c.SendStatus(204)
}
