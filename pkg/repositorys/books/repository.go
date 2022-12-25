package books

import (
	"gobackend/pkg/entities"
)

type Repository interface {
	CreateBook() (*entities.Books, error)
}

// type repository struct {
// 	Db *database.DB
// }
