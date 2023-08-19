package handlers

import (
	"errors"
	"gobackend/database"
	"gobackend/pkg"
	"gobackend/pkg/auth"
	"gobackend/pkg/entities"
	"log"
	"os"
	"time"

	validate "gobackend/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

var (
	Googleoauth *oauth2.Config
)

func Googleapi() {
	Googleoauth = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		RedirectURL:  "https://api.pilput.dev/auth/oauth/callback",
		Endpoint:     google.Endpoint,
	}
}
func Loginoatuth(c *fiber.Ctx) error {
	// Create oauthState cookie
	url := Googleoauth.AuthCodeURL("state")
	return c.Redirect(url)
}

func CallbackHandler(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		code := c.Query("code")
		token, err := Googleoauth.Exchange(c.Context(), code)
		if err != nil {
			log.Println("Failed to exchange token:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token")
		}
		profile, err := service.GetUserInfoGoogle(token.AccessToken)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		user, err := service.GetUserOrCreate(profile.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		jwttoken, err := service.SetTokenJwt(user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		domain := os.Getenv("DOMAIN")
		cookie := fiber.Cookie{
			Name:     "token",
			Domain:   "." + domain,
			Value:    jwttoken,
			Expires:  time.Now().Add(time.Hour * 24),
			SameSite: "strict",
		}
		c.Cookie(&cookie)
		return c.Redirect("https://" + domain)
	}
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(c *fiber.Ctx) error {
	db := database.DB
	type LoginInput struct {
		Email    string `json:"Email" validate:"required"`
		Password string `jsno:"password" validate:"required,min=8" `
	}

	var user entities.Users

	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	resulvalidate := validate.ValidateThis(input)
	if resulvalidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(resulvalidate)
	}

	email := input.Email
	err := db.Where("email = ?", email).First(&user).Error
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
	claims["id"] = user.ID
	claims["email"] = email
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
