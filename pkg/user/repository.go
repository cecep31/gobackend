package user

import (
	"gobackend/pkg/entities"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(uer *Users) (*Users, error)
	GetUsers() (*[]entities.Users, error)
	GetUser(user *entities.Users) (*entities.Users, error)
	CreateUserWithOutValidate(user *entities.Users) (*entities.Users, error)
}

type repository struct {
	Db *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repository{
		Db: db,
	}
}

func (r *repository) CreateUser(user *Users) (*Users, error) {
	err := r.Db.Create(&user).Error
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}

}
func (r *repository) CreateUserWithOutValidate(user *entities.Users) (*entities.Users, error) {
	err := r.Db.Create(&user).Error
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}

}

func (r *repository) GetUsers() (*[]entities.Users, error) {
	var users []entities.Users
	result := r.Db.Find(&users)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *repository) GetUser(user *entities.Users) (*entities.Users, error) {
	err := r.Db.First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
