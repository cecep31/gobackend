package handlers

import (
	"fmt"
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
		Scopes:       []string{"email", "profile"},
		RedirectURL:  "https://api.pilput.dev/oauth/callback",
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

	token, err := Googleoauth.Exchange(c.Context(), code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token)

	// Lakukan sesuatu dengan token akses

	return c.SendString("tok")
}
