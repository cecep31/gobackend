package user

import (
	"gobackend/pkg/entities"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(uer *Users) (*Users, error)
	GetUsers() (*[]entities.User, error)
	GetUser(user *entities.User) (*entities.User, error)
	CreateUserWithOutValidate(user *entities.User) (*entities.User, error)
	UpdateUser(user *entities.User) error
	DeleteUser(user *entities.User) error
	GetUserByEmail(email string) (*entities.User, error)
	GetWriter() (interface{}, error)
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
func (r *repository) CreateUserWithOutValidate(user *entities.User) (*entities.User, error) {
	err := r.Db.Create(&user).Error
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}

}

func (r *repository) GetUsers() (*[]entities.User, error) {
	var users []entities.User
	result := r.Db.Find(&users)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *repository) GetUser(user *entities.User) (*entities.User, error) {
	err := r.Db.First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) UpdateUser(user *entities.User) error {
	err := r.Db.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetWriter() (interface{}, error) {
	result := []map[string]interface{}{}
	err := r.Db.Table("users").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (r *repository) DeleteUser(user *entities.User) error {
	err := r.Db.Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}
