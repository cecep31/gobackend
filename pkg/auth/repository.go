package auth

import (
	"gobackend/pkg/entities"

	"gorm.io/gorm"
)

type Repository interface {
	GetUserByEmail(email string) (*entities.Users, error)
	UpdateUser(user *entities.Users) error
}

type repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		Db: db,
	}
}

func (r *repository) GetUserByEmail(email string) (*entities.Users, error) {
	user := new(entities.Users)
	if err := r.Db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (repo *repository) UpdateUser(user *entities.Users) error {
	return repo.Db.Save(user).Error
}
