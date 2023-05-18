package user

import (
	"gobackend/pkg/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(uer *entities.Users) (*entities.Users, error)
	GetUsers() (*[]entities.Users, error)
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
	user.ID = uuid.New()
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