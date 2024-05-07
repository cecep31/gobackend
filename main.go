package main

import (
	"fmt"
	"gobackend/pkg/auth"
	"gobackend/pkg/utils"
	"log"

	"gobackend/api"
	"gobackend/api/handlers"
	"gobackend/database"
	"gobackend/pkg/posts"
	"gobackend/pkg/storage"
	"gobackend/pkg/user"
	"gobackend/server"
	initstorage "gobackend/storage"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("dotenv not found")
	}

	database.SetupDatabase()
	db := database.DB
	minio, errstore := initstorage.NewFileStorageClient()
	if errstore != nil {
		log.Fatal(errstore)
	}

	database.MigrationDB()

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
	fmt.Println("Initial repository & service Done")

	app := server.Create()

	v2 := app.Group("api/v2")
	auth := app.Group("auth")
	api.SetupAuthRoutes(auth, authservice)
	api.UserRouter(v2, userserivce)
	api.PostRouter(v2, postservice)

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}
}
