package auth

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `jsno:"password"`
	}
	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	identity := input.Identity
	pass := input.Password
	if identity != "pilput" || pass != "pilput" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	role := "admin"
	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = identity
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("SIGNKEY")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	var data = map[string]string{
		"name":  identity,
		"role":  role,
		"token": t,
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success login",
		"data":    data,
	})

}
