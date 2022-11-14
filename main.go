package main

import (
	"log"

	"github.com/cecep31/gobackend/api"
	"github.com/cecep31/gobackend/api/books"
	"github.com/cecep31/gobackend/api/items"
	"github.com/cecep31/gobackend/database"
	"github.com/cecep31/gobackend/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Server initialization
	app := server.Create()

	// Migrations
	database.DB.AutoMigrate(&books.Book{}, &items.Items{})

	// Api routes
	api.Setup(app)

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}
}
