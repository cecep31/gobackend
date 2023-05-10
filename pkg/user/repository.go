package user

import (
	"gobackend/pkg/entities"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(uer *entities.Users) (*entities.Users, error)
}

type repository struct {
	Db *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repository{
		Db: db,
	}
}

func (r *repository) CreateUser(user *entities.Users) (*entities.Users, error) {
	err := r.Db.Create(user).Error
	if err != nil {
		return user, nil
	} else {
		return nil, err
	}

}
