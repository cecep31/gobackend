package main

import (
	"fmt"
	"gobackend/pkg/auth"
	"gobackend/pkg/tasks"
	"gobackend/pkg/utils"
	"log"
	"os"

	"gobackend/api"
	"gobackend/api/handlers"
	"gobackend/database"
	"gobackend/pkg/entities"
	"gobackend/pkg/posts"
	"gobackend/pkg/storage"
	"gobackend/pkg/user"
	"gobackend/server"
	initstorage "gobackend/storage"
	"gobackend/ws"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("dotenv not found")
	}

	db := database.SetupDatabase()
	minio := initstorage.InitFileStorage()

	if os.Getenv("MIGRATE") != "" {
		fmt.Println("Migration...")
		db.AutoMigrate(&entities.Users{}, &entities.Tasks{}, &entities.Posts{}, &entities.PostComments{})
		fmt.Println("Migration Done")
	}

	handlers.Googleapi()
	utils.SetupValidate()

	fmt.Println("Initial repository & service")
	miniorepo := storage.NewRepo(minio)
	userrepo := user.NewRepo(db)
	userserivce := user.NewService(userrepo)
	postrepo := posts.NewRepo(db)
	postservice := posts.NewService(postrepo, miniorepo)
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

	ws.WsSetup(app)

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}
}
