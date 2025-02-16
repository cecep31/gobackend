package handlers

import (
	"errors"
	"gobackend/api/presenter"
	"gobackend/pkg"
	"gobackend/pkg/auth"
	"gobackend/pkg/entities"
	"gobackend/pkg/utils"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

func LoginWithOAuth(ctx *fiber.Ctx) error {
	authURL := Googleoauth.AuthCodeURL("state")
	return ctx.Redirect(authURL)
}

func HandleOAuthCallback(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		code := c.FormValue("code")
		token, err := Googleoauth.Exchange(c.Context(), code)
		if err != nil {
			log.Println("Failed to exchange token:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token" + err.Error())
		}
		profile, err := service.GetUserInfoGoogle(token.AccessToken)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		user, err := service.GetUserOrCreate(profile)
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

func HandleLogin(authservice auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var logininput auth.LoginInput

		if err := c.BodyParser(&logininput); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		resulvalidate := utils.ValidateThis(logininput)

		if len(resulvalidate) > 0 {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ErrorResponse(resulvalidate))
		}

		user, err := authservice.GetUserByEmail(logininput.Email)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid email or password", "data": nil})
			} else {
				return pkg.Unexpected(err.Error())
			}
		}

		if !utils.CheckPasswordHash(logininput.Password, user.Password) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid email or password", "data": nil})
		}

		t, err := authservice.GenerateToken(user)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// c.Cookie(&fiber.Cookie{
		// 	Name:     "token",
		// 	Value:    t,
		// 	Domain:   "." + os.Getenv("DOMAIN"),
		// 	Path:     "/",
		// 	Expires:  time.Now().Add(time.Hour * 168),
		// 	Secure:   false,
		// 	HTTPOnly: true,
		// })

		return c.JSON(fiber.Map{
			"access_token": t,
		})

	}
}

func GetUserProfile(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userlocal := c.Locals("user").(*jwt.Token)
		claims := userlocal.Claims.(jwt.MapClaims)
		id := claims["id"].(string)
		profile, err := service.GetProfile(id)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(presenter.UserSuccessResponse(profile))
	}
}

func UpdateUserProfile(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userlocal := c.Locals("user").(*jwt.Token)
		claims := userlocal.Claims.(jwt.MapClaims)
		id := claims["id"].(string)

		var requestBody entities.User

		if err := c.BodyParser(&requestBody); err != nil {
			return err
		}

		resulvalidate := utils.ValidateThis(requestBody)
		if len(resulvalidate) > 0 {
			return c.JSON(presenter.ErrorResponse(resulvalidate))
		}

		user, err := service.GetProfile(id)
		if err != nil {
			return c.Status(404).JSON(presenter.ErrorResponse(err))
		}

		user.FirstName = requestBody.FirstName
		user.LastName = requestBody.LastName
		user.Image = requestBody.Image

		error := service.UpdateProfile(user)
		if error != nil {
			return c.Status(404).JSON(presenter.ErrorResponse(error))
		}
		return c.Status(200).JSON(user)
	}
}
