package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"gobackend/database"
	"gobackend/pkg"
	"gobackend/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users struct {
	database.DefaultModel
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Image    string `json:"image" gorm:"type:text"`
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	fmt.Println(hash)
	fmt.Println(err)
	return err == nil

}

func Login(c *fiber.Ctx) error {
	db := database.DB
	type LoginInput struct {
		Username string `json:"username"`
		Password string `jsno:"password"`
	}
	var user entities.Users

	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	username := input.Username
	err := db.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	} else if err != nil {
		return pkg.Unexpected(err.Error())
	}

	pass := input.Password

	if !CheckPasswordHash(pass, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["issuperadmin"] = user.Issuperadmin
	claims["exp"] = time.Now().Add(time.Hour * 168).Unix()

	t, err := token.SignedString([]byte(os.Getenv("SIGNKEY")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"access_token": t,
	})

}
