package posts

import (
	"fmt"
	"gobackend/pkg/entities"
	"gobackend/pkg/storage"
	"gobackend/pkg/utils"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

type Service interface {
	InserPosts(post *entities.Post) (*entities.Post, error)
	GetPosts() (*[]entities.Post, error)
	GetPost(slug string) (*entities.Post, error)
	GetPostsRandom() (*[]entities.Post, error)
	GetTotalPosts() (int64, error)
	GetPostsPaginated(offset int, Limit int) ([]entities.Post, error)
	UpdatePost(post *PostUpdate) error
	GetPostByid(id string) (*entities.Post, error)
	DeletePost(id string) error
	UploadPhoto(ctx *fasthttp.RequestCtx, file *multipart.FileHeader, uploader uuid.UUID) (string, error)
	ValidFileExtension(filename string, allowedExtensions []string) bool
}

type service struct {
	repository      Repository
	minioRepository storage.Repository
}

func NewService(r Repository, miniorepo storage.Repository) Service {
	return &service{
		repository:      r,
		minioRepository: miniorepo,
	}
}

func (s *service) UpdatePost(post *PostUpdate) error {
	newpost := entities.Post{Title: post.Title, Body: post.Body, CreatedBy: post.CreatedBy, Slug: post.Slug, Photo_url: post.Photo_url}
	return s.repository.UpdatePost(&newpost)
}
func (s *service) InserPosts(post *entities.Post) (*entities.Post, error) {
	return s.repository.CreatePost(post)
}
func (s *service) GetPosts() (*[]entities.Post, error) {
	var posts []entities.Post
	return s.repository.GetPosts(&posts)
}

func (s *service) GetPostsRandom() (*[]entities.Post, error) {
	return s.repository.GetPostsRandom(6)
}

func (s *service) GetPost(slug string) (*entities.Post, error) {
	return s.repository.GetPostBySlug(slug)
}
func (s *service) GetPostByid(id string) (*entities.Post, error) {
	post := new(entities.Post)
	id_uuid, _ := uuid.Parse(id)
	post.ID = id_uuid
	return s.repository.GetPost(id, post)
}

func (s *service) GetTotalPosts() (int64, error) {
	return s.repository.CountPosts()
}

func (s *service) GetPostsPaginated(offset int, Limit int) ([]entities.Post, error) {
	return s.repository.FindPaginated(offset, Limit)
}
func (s *service) DeletePost(id string) error {
	id_uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	post := new(entities.Post)
	post.ID = id_uuid
	return s.repository.DeletePostById(post)
}

func (s *service) UploadPhoto(ctx *fasthttp.RequestCtx, file *multipart.FileHeader, uploader uuid.UUID) (string, error) {
	generatedfilename, _ := utils.GenerateRandomString(15, file.Filename)
	uploadFilename := fmt.Sprintf("post_photo/%s", generatedfilename)
	if err := s.minioRepository.UploadFile(ctx, uploadFilename, file); err != nil {
		return "", err
	}
	if s.minioRepository.AddFileRecord(file.Filename, uploadFilename, file.Size, file.Header.Get("Content-Type"), uploader) != nil {
		return "/" + uploadFilename, nil
	}
	return "/" + uploadFilename, nil
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
