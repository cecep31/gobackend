package posts

import (
	"gobackend/pkg/entities"

	"github.com/google/uuid"
)

type Service interface {
	InserPosts(user *entities.Posts) (*entities.Posts, error)
	GetPosts() (*[]entities.Posts, error)
	GetPost(id uuid.UUID) (*entities.Posts, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InserPosts(user *entities.Posts) (*entities.Posts, error) {
	return s.repository.CreatePost(user)
}
func (s *service) GetPosts() (*[]entities.Posts, error) {
	return s.repository.GetPosts()
}
func (s *service) GetPost(id uuid.UUID) (*entities.Posts, error) {
	var post entities.Posts
	post.ID = id
	return s.repository.GetPost(&post)
}
