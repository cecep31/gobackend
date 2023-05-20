package handlers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	Googleoauth *oauth2.Config
)

func Googleapi() {
	Googleoauth = &oauth2.Config{
		ClientID:     "10345906756-nqkqh6o3k5ea7vbr4b4khee9ivjnf59f.apps.googleusercontent.com",
		ClientSecret: "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",
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
