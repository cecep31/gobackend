package user

import (
	"gobackend/pkg/entities"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	InserUser(user *Users) (*Users, error)
	GetUsers() (*[]entities.Users, error)
	GetUser(id uuid.UUID) (*entities.Users, error)
	DeleteUser(user *entities.Users) error
	GetWriter() (interface{}, error)
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
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func (s *service) InserUser(user *Users) (*Users, error) {
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
func (s *service) GetUser(id uuid.UUID) (*entities.Users, error) {
	var user entities.Users
	user.ID = id
	return s.repository.GetUser(&user)
}
func (s *service) DeleteUser(user *entities.Users) error {
	return s.repository.DeleteUser(user)
}

func (s *service) GetWriter() (interface{}, error) {
	return s.repository.GetWriter()
}
