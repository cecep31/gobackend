package main

import (
	"gobackend/pkg/auth"
	"gobackend/pkg/tasks"
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

	db := database.SetupDatabase()
	if os.Getenv("MIGRATE") != "" {
		println("Migration...")
		db.AutoMigrate(&entities.Items{}, &entities.Users{}, &entities.Tasks{}, &entities.Taskgorups{}, &entities.Posts{})
	}

	handlers.Googleapi()
	storage.InitFileStorage()

	userrepo := user.NewRepo(db)
	userserivce := user.NewService(userrepo)
	postrepo := posts.NewRepo(db)
	postservice := posts.NewService(postrepo)
	authrepo := auth.NewRepository(db)
	authservice := auth.NewService(authrepo)
	taskrepo := tasks.NewRepository(db)
	taskservice := tasks.NewService(taskrepo)
	app := server.Create()

	v2 := app.Group("api/v2")
	api.UserRouter(v2, userserivce)
	api.AuthRouter(v2, authservice)
	api.PostRouter(v2, postservice)
	api.TaskRouter(v2, taskservice)

	// Api routes
	api.Setup(app)
	ws.WsSetup(app)

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}
}
