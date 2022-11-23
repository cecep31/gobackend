package users

import (
	"errors"

	"github.com/cecep31/gobackend/database"
	"github.com/cecep31/gobackend/pkg"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	database.DefaultModel
	Username string `json:"username"`
	Password string `json:"password" gorm:"type:text"`
	Image    string `json:"image" gorm:"type:text"`
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []User
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
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return pkg.BadRequest("Invalid params")
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

	var user User
	err := db.First(&user, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkg.EntityNotFound("No user found")
	} else if err != nil {
		return pkg.Unexpected(err.Error())
	}
	db.Delete(&user)
	return c.SendStatus(204)
}
