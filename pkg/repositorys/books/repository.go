package books

import (
	"gobackend/pkg/entities"
)

type Repository interface {
	CreateBook() (*entities.Book, error)
}

// type repository struct {
// 	Db *database.DB
// }
