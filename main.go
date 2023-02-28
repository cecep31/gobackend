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
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println("Error loading .env file")
	// }
	godotenv.Load()
	// Server initialization
	app := server.Create()

	// Migrations
	println("Migration...")
	database.DB.AutoMigrate(&entities.Books{}, &entities.Items{}, &entities.Users{}, &entities.Tasks{}, &entities.Taskgorups{}, &entities.Posts{}, &entities.Posttags{})

	// Api routes
	api.Setup(app)

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}
}
