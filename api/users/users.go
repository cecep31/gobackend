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
	var users []entities.Users
	db.Select("id", "username", "role", "email", "issuperadmin", "image").Find(&users)
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data":    users,
	})
}

func Getuser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user entities.Users
	err := db.First(&user, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkg.EntityNotFound("record Not Found")
	} else if err != nil {
		return pkg.Unexpected(err.Error())
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data":    user,
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func NewUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(entities.Users)
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

	var user entities.Users
	err := db.First(&user, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkg.EntityNotFound("No user found")
	} else if err != nil {
		return pkg.Unexpected(err.Error())
	}
	db.Delete(&user)
	return c.SendStatus(204)
}

func UpdateUser(c *fiber.Ctx) error {
	type uservalidate struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}
	validate := new(uservalidate)
	var user entities.Users

	db := database.DB
	id := c.Params("id")
	err := db.First(&user, id).Error
	// return c.JSON(user)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkg.EntityNotFound("No book found")
	} else if err != nil {
		// return pkg.Unexpected(err.Error())
		return pkg.Unexpected("i dont know")
	}

	if err := c.BodyParser(validate); err != nil {
		return pkg.BadRequest("invalid params")
	}

	db.Model(&user).Updates(&entities.Users{Email: validate.Email, Username: validate.Username})

	return c.JSON(validate)

}
