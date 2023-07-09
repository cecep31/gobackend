package tasks

import (
	"gobackend/pkg/entities"
)

type Service interface {
	GetTasks() (*[]entities.Tasks, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetTasks() (*[]entities.Tasks, error) {
	return s.repository.GetTasks()
}
