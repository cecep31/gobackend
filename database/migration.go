package database

// import (
// 	"fmt"
// 	"gobackend/pkg/entities"
// 	"os"

// 	"gorm.io/gorm"
// )

// func MigrationDB(db *gorm.DB) {
// 	if os.Getenv("MIGRATE") != "" {
// 		fmt.Println("Migration...")
// 		db.AutoMigrate(&entities.Users{}, &entities.Tasks{}, &entities.Posts{}, &entities.PostComments{})
// 		fmt.Println("Migration Done")
// 	}
// }
