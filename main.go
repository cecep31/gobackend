package main

import (
	"log"

	"gobackend/api"
	"gobackend/database"
	"gobackend/pkg/user"
	"gobackend/server"
	"gobackend/storage"
	"gobackend/ws"

	"github.com/joho/godotenv"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println("Error loading .env file")
	// }
	godotenv.Load()
	// Server initialization
	database.SetupDatabase()
	storage.InitFileStorage()
	userrepo := user.NewRepo(database.DB)
	userserivce := user.NewService(userrepo)
	app := server.Create()

	v2 := app.Group("api/v2")
	api.UserRouter(v2, userserivce)

	// Migrations
	println("Migration...")
	// database.DB.AutoMigrate(&entities.Books{}, &entities.Items{}, &entities.Users{}, &entities.Tasks{}, &entities.Taskgorups{}, &entities.Posts{}, &entities.Posttags{}, &entities.Globalchat{})

	// Api routes
	api.Setup(app)
	ws.WsSetup(app)

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}
}
