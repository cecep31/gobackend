package posts

import (
	"gobackend/pkg/entities"
	"gobackend/pkg/storage"
	"gobackend/pkg/utils"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

type Service interface {
	InserPosts(post *entities.Posts) (*entities.Posts, error)
	GetPosts() (*[]entities.Posts, error)
	GetPost(slug string) (*entities.Posts, error)
	GetPostsRandom() (*[]entities.Posts, error)
	GetTotalPosts() (int64, error)
	GetPostsPaginated(page int, perPage int) ([]entities.Posts, error)
	UpdatePost(post *PostUpdate) error
	GetPostByid(id string) (*entities.Posts, error)
	DeletePost(id string) error
	PutObjectPhoto(ctx *fasthttp.RequestCtx, objectname string, file *multipart.FileHeader) (string, error)
	ValidFileExtension(filename string, allowedExtensions []string) bool
}

type service struct {
	repository Repository
	miniorepo  storage.Repository
}

func NewService(r Repository, miniorepo storage.Repository) Service {
	return &service{
		repository: r,
		miniorepo:  miniorepo,
	}
}

func (s *service) UpdatePost(post *PostUpdate) error {
	return s.repository.UpdatePost(post)
}
func (s *service) InserPosts(post *entities.Posts) (*entities.Posts, error) {
	return s.repository.CreatePost(post)
}
func (s *service) GetPosts() (*[]entities.Posts, error) {
	var posts []entities.Posts
	return s.repository.GetPosts(&posts)
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

func (s *service) PutObjectPhoto(ctx *fasthttp.RequestCtx, objectname string, file *multipart.FileHeader) (string, error) {

	newobjcetname := "post_photo/" + utils.GenerateRandomFilename(objectname)
	err := s.miniorepo.UploadFile(ctx, newobjcetname, file)
	if err != nil {
		return "", err
	}
	return newobjcetname, nil
}

func (s *service) ValidFileExtension(filename string, allowedExtensions []string) bool {
	ext := filepath.Ext(filename)
	for _, validExt := range allowedExtensions {
		if ext == validExt {
			return true
		}
	}
	return false
}
