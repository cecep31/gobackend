package posts

import (
	"gobackend/pkg/entities"

	"github.com/google/uuid"
)

type Service interface {
	InserPosts(post *entities.Posts) (*entities.Posts, error)
	GetPosts() (*[]entities.Posts, error)
	GetPost(slug string) (*entities.Posts, error)
	GetPostsRandom() (*[]entities.Posts, error)
	GetTotalPosts() (int64, error)
	GetPostsPaginated(page int, perPage int) ([]entities.Posts, error)
	UpdatePost(post *Posts) error
	GetPostByid(id string) (*entities.Posts, error)
	DeletePost(id string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) UpdatePost(post *Posts) error {
	return s.repository.UpdatePost(post)
}
func (s *service) InserPosts(post *entities.Posts) (*entities.Posts, error) {
	return s.repository.CreatePost(post)
}
func (s *service) GetPosts() (*[]entities.Posts, error) {
	return s.repository.GetPosts()
}

func (s *service) GetPostsRandom() (*[]entities.Posts, error) {
	return s.repository.GetPostsRandom(6)
}

func (s *service) GetPost(slug string) (*entities.Posts, error) {
	return s.repository.GetPostBySlug(slug)
}
func (s *service) GetPostByid(id string) (*entities.Posts, error) {
	post := new(entities.Posts)
	id_uuid, _ := uuid.Parse(id)
	post.ID = id_uuid
	return s.repository.GetPost(id, post)
}

func (s *service) GetTotalPosts() (int64, error) {
	return s.repository.Count()
}

func (s *service) GetPostsPaginated(page int, perPage int) ([]entities.Posts, error) {
	return s.repository.FindPaginated(page, perPage)
}
func (s *service) DeletePost(id string) error {
	id_uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	post := new(entities.Posts)
	post.ID = id_uuid

	return s.repository.DeletePostById(post)
}
