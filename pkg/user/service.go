package user

import (
	"gobackend/pkg/entities"
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

func (s *service) InserUser(user *entities.Users) (*entities.Users, error) {
	return s.repository.CreateUser(user)
}
func (s *service) GetUsers() (*[]entities.Users, error) {
	return s.repository.GetUsers()
}
