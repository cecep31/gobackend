package handlers

import (
	"fmt"
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
		Scopes:       []string{"email", "profile"},
		RedirectURL:  "https://api.pilput.dev/api/v2/oauth/callback",
		Endpoint:     google.Endpoint,
	}
}
func Loginoatuth(c *fiber.Ctx) error {
	// Create oauthState cookie
	url := Googleoauth.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

func CallbackHandler(c *fiber.Ctx) error {
	code := c.Query("code")
	fmt.Println(code)

	token, err := Googleoauth.Exchange(c.Context(), code)
	if err != nil {
		fmt.Println(err.Error())

	}
	fmt.Println(token)

	// Lakukan sesuatu dengan token akses

	return c.SendString("ok")
}
