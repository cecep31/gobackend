package user

import (
	"gobackend/pkg/entities"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	InserUser(user *entities.Users) (*entities.Users, error)
	GetUsers() (*[]entities.Users, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *service) InserUser(user *entities.Users) (*entities.Users, error) {
	pass := user.Password
	hashpass, err := HashPassword(pass)
	if err != nil {
		return nil, err
	}
	user.Password = hashpass
	return s.repository.CreateUser(user)
}
func (s *service) GetUsers() (*[]entities.Users, error) {
	return s.repository.GetUsers()
}
