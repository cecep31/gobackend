package database

import (
	"fmt"
	"gobackend/pkg/entities"
	"os"
)

func MigrationDB() {
	if os.Getenv("MIGRATE") != "" {
		fmt.Println("Migration...")
		DB.AutoMigrate(&entities.Users{}, &entities.Tasks{}, &entities.Posts{}, &entities.PostComments{})
		fmt.Println("Migration Done")
	}
}
