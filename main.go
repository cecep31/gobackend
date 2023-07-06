package main

import (
	"gobackend/pkg/auth"
	"log"
	"os"

	"gobackend/api"
	"gobackend/api/handlers"
	"gobackend/database"
	"gobackend/pkg/entities"
	"gobackend/pkg/posts"
	"gobackend/pkg/user"
	"gobackend/server"
	"gobackend/storage"
	"gobackend/ws"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// Server initialization
	db := database.SetupDatabase()
	if os.Getenv("DEBUG") != "" {
		println("Migration...")
		db.AutoMigrate(&entities.Items{}, &entities.Users{}, &entities.Tasks{}, &entities.Taskgorups{}, &entities.Posts{}, &entities.Posttags{}, &entities.Globalchat{})
	}

	handlers.Googleapi()

	storage.InitFileStorage()

	userrepo := user.NewRepo(db)
	userserivce := user.NewService(userrepo)
	postrepo := posts.NewRepo(db)
	postservice := posts.NewService(postrepo)
	authrepo := auth.NewRepository(db)
	authservice := auth.NewService(authrepo)
	app := server.Create()

	v2 := app.Group("api/v2")
	api.UserRouter(v2, userserivce)
	api.AuthRouter(v2, authservice)
	api.PostRouter(v2, postservice)

	// Api routes
	api.Setup(app)
	ws.WsSetup(app)

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}
}
