package auth

import (
	"gobackend/pkg/entities"

	"gorm.io/gorm"
)

type Repository interface {
	GetUserByEmail(email string) (*entities.User, error)
	UpdateUser(user *entities.User) error
}

type repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		Db: db,
	}
}

func (r *repository) GetUserByEmail(email string) (*entities.User, error) {
	user := new(entities.User)
	if err := r.Db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (repo *repository) UpdateUser(user *entities.User) error {
	return repo.Db.Save(user).Error
}
