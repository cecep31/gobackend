package handlers

import (
	"fmt"
	"gobackend/pkg/auth"
	"io/ioutil"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	Googleoauth *oauth2.Config
)

func Googleapi() {
	Googleoauth = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		RedirectURL:  "http://localhost:8080/api/v2/oauth/callback",
		Endpoint:     google.Endpoint,
	}
}
func Loginoatuth(c *fiber.Ctx) error {
	// Create oauthState cookie
	url := Googleoauth.AuthCodeURL("state")
	fmt.Print(url)
	return c.Redirect(url)
}

func CallbackHandler(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Print("masuk callback")
		code := c.Query("code")
		//		fmt.Println(code)

		token, err := Googleoauth.Exchange(c.Context(), code)
		if err != nil {
			log.Println("Failed to exchange token:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token")
		}
		profile := service.GetUserInfoGoogle(token.AccessToken)
		return c.JSON(profile)
	}
}
