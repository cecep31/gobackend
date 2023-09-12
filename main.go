package main

import (
	"fmt"
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

	validate "gobackend/pkg/validator"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := database.SetupDatabase()
	if os.Getenv("MIGRATE") != "" {
		fmt.Println("Migration...")
		db.AutoMigrate(&entities.Users{}, &entities.Tasks{}, &entities.Posts{}, &entities.PostComments{})
		fmt.Println("Migration Done")
	}

	handlers.Googleapi()
	storage.InitFileStorage()
	validate.SetupValidate()

	fmt.Println("Initial repository & service")
	userrepo := user.NewRepo(db)
	userserivce := user.NewService(userrepo)
	postrepo := posts.NewRepo(db)
	postservice := posts.NewService(postrepo)
	authrepo := auth.NewRepository(db)
	authservice := auth.NewService(authrepo, userrepo)
	taskrepo := tasks.NewRepository(db)
	taskservice := tasks.NewService(taskrepo)
	fmt.Println("Initial repository & service Done")
	app := server.Create()

	v2 := app.Group("api/v2")
	auth := app.Group("auth")
	api.AuthRouter(auth, authservice)
	api.UserRouter(v2, userserivce)
	api.PostRouter(v2, postservice)
	api.TaskRouter(v2, taskservice)

	// Api routes
	ws.WsSetup(app)

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}
}
