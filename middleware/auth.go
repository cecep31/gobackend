package middleware

import (
	"gobackend/database"
	"gobackend/pkg"
	"gobackend/pkg/entities"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("SIGNKEY")),
		ErrorHandler: jwtError,
	})
}

// func ProtectedSuperAdmin() func(*fiber.Ctx) error {
// 	return jwtware.New(jwtware.Config{
// 		SigningKey:   []byte(os.Getenv("SIGNKEY")),
// 		ErrorHandler: jwtErrorSuperAdmin,
// 	})
// }

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
	}
}

func IsSuperAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	// println(c.Locals("user").(*jwt.Token))
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	if role != "super_admin" {
		return pkg.CredentionProtect("Need as super_admin")
	} else {
		return c.Next()
	}
}

func GetUser(c *fiber.Ctx) error {
	userlocal := c.Locals("user").(*jwt.Token)
	claims := userlocal.Claims.(jwt.MapClaims)
	username := claims["identity"].(string)
	db := database.DB
	var userdata entities.User
	db.Where("username = ?", username).First(&userdata)

	c.Locals("datauser", userdata)
	return c.Next()
}
