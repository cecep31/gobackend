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
	var user entities.Users
	err := r.Db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *repository) UpdateUser(user *entities.Users) error {
	err := r.Db.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}
