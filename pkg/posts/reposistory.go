package posts

import (
	"gobackend/pkg/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreatePost(post *entities.Posts) (*entities.Posts, error)
	GetPosts() (*[]entities.Posts, error)
	GetPost(user *entities.Posts) (*entities.Posts, error)
	GetPostBySlug(slug string) (*entities.Posts, error)
	GetPostsRandom(take int) (*[]entities.Posts, error)
	Count() (int64, error)
	FindPaginated(page int, perPage int) ([]entities.Posts, error)
	UpdatePost(post *Posts) error
}

type repository struct {
	Db *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repository{
		Db: db,
	}
}

func (r *repository) CreatePost(post *entities.Posts) (*entities.Posts, error) {
	post.ID = uuid.New()
	err := r.Db.Create(&post).Error
	if err != nil {
		return nil, err
	} else {
		return post, nil
	}

}
func (r *repository) UpdatePost(post *Posts) error {
	return r.Db.Model(post).Updates(entities.Posts{Title: post.Title, Slug: post.Slug, Body: post.Body}).Error
}
func (r *repository) GetPosts() (*[]entities.Posts, error) {
	var posts []entities.Posts
	result := r.Db.Preload("Creator").Order("created_at DESC").Find(&posts)
	err := result.Error
	if err != nil {
		return &[]entities.Posts{}, err
	}
	return &posts, nil
}

func (r *repository) FindPaginated(page int, perPage int) ([]entities.Posts, error) {
	var posts []entities.Posts
	offset := (page - 1) * perPage
	result := r.Db.Offset(offset).Preload("Creator").Order("created_at desc").Limit(perPage).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (r *repository) GetPostsRandom(take int) (*[]entities.Posts, error) {
	posts := new([]entities.Posts)
	result := r.Db.Preload("Creator").Order("RANDOM()").Limit(take).Find(posts)
	err := result.Error
	if err != nil {
		return &[]entities.Posts{}, err
	}
	return posts, nil
}

func (r *repository) GetPost(post *entities.Posts) (*entities.Posts, error) {
	err := r.Db.First(post).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *repository) GetPostBySlug(slug string) (*entities.Posts, error) {
	post := new(entities.Posts)
	err := r.Db.First(post, "slug = ?", slug).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *repository) Count() (int64, error) {
	var count int64
	result := r.Db.Model(&entities.Posts{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
