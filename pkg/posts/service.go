package posts

import (
	"gobackend/pkg/entities"
)

type Service interface {
	InserPosts(user *entities.Posts) (*entities.Posts, error)
	GetPosts() (*[]entities.Posts, error)
	GetPost(slug string) (*entities.Posts, error)
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

func (s *service) GetPost(slug string) (*entities.Posts, error) {
	return s.repository.GetPostBySlug(slug)
}
