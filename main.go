package main

import (
	"log"

	"gobackend/api"
	"gobackend/database"
	"gobackend/pkg/entities"
	"gobackend/server"

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
	database.DB.AutoMigrate(&entities.Book{}, &entities.Items{}, &entities.User{})

	// Api routes
	api.Setup(app)

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}
}
