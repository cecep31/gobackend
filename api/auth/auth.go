package auth

import (
	"errors"
	"os"
	"time"

	"github.com/cecep31/gobackend/api/users"
	"github.com/cecep31/gobackend/database"
	"github.com/cecep31/gobackend/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(c *fiber.Ctx) error {
	db := database.DB
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `jsno:"password"`
	}
	var user users.User

	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	identity := input.Identity
	err := db.Where("username = ?", identity).First(&user).Error

	// check ada error di query
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.SendStatus(fiber.StatusUnauthorized)
	} else if err != nil {
		return pkg.Unexpected(err.Error())
	}

	username := user.Password
	println(username)
	pass := input.Password

	if !CheckPasswordHash(pass, user.Password) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = identity
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("SIGNKEY")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success login",
		"data":    t,
	})

}
