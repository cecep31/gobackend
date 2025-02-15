package posts

import (
	"gobackend/pkg/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreatePost(post *entities.Post) (*entities.Post, error)
	GetPosts(posts *[]entities.Post) (*[]entities.Post, error)
	GetPost(id string, post *entities.Post) (*entities.Post, error)
	GetPostBySlug(slug string) (*entities.Post, error)
	GetPostsRandom(take int) (*[]entities.Post, error)
	CountPosts() (int64, error)
	FindPaginated(offset int, Limit int) ([]entities.Post, error)
	UpdatePost(post *entities.Post) error
	DeletePostById(post *entities.Post) error
}

type repository struct {
	Db *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repository{
		Db: db,
	}
}

func (r *repository) CreatePost(post *entities.Post) (*entities.Post, error) {
	post.ID = uuid.New()
	err := r.Db.Create(&post).Error
	if err != nil {
		return nil, err
	} else {
		return post, nil
	}
}

func (r *repository) UpdatePost(post *entities.Post) error {
	return r.Db.Model(post).Updates(entities.Post{Title: post.Title, Slug: post.Slug, Body: post.Body}).Error
}

func (r *repository) GetPosts(posts *[]entities.Post) (*[]entities.Post, error) {
	result := r.Db.Preload("Creator").Order("created_at DESC").Find(posts)
	err := result.Error
	if err != nil {
		return &[]entities.Post{}, err
	}
	return posts, nil
}

func (r *repository) FindPaginated(offset int, Limit int) ([]entities.Post, error) {
	var posts []entities.Post
	result := r.Db.Offset(offset).Preload("Tags").Preload("Creator").Preload("Likes").Order("created_at desc").Limit(Limit).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (r *repository) GetPostsRandom(take int) (*[]entities.Post, error) {
	posts := new([]entities.Post)
	result := r.Db.Preload("Creator").Preload("Tags").Order("RANDOM()").Limit(take).Find(posts)
	err := result.Error
	if err != nil {
		return &[]entities.Post{}, err
	}
	return posts, nil
}

func (r *repository) GetPost(id string, post *entities.Post) (*entities.Post, error) {
	err := r.Db.Preload("Creator").Preload("Tags").Preload("Likes").Where("id = ?", id).First(post).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (repo *repository) GetPostBySlug(slug string) (*entities.Post, error) {
	post := new(entities.Post)
	if err := repo.Db.Preload("Creator").Preload("Tags").Preload("Likes").Where("slug = ?", slug).First(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *repository) DeletePostById(post *entities.Post) error {
	err := r.Db.Delete(post).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CountPosts() (int64, error) {
	var postCount int64
	result := r.Db.Model(&entities.Post{}).Count(&postCount)
	return postCount, result.Error
}
