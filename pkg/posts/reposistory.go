package posts

import (
	"gobackend/pkg/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreatePost(post *entities.Posts) (*entities.Posts, error)
	GetPosts(posts *[]entities.Posts) (*[]entities.Posts, error)
	GetPost(id string, post *entities.Posts) (*entities.Posts, error)
	GetPostBySlug(slug string) (*entities.Posts, error)
	GetPostsRandom(take int) (*[]entities.Posts, error)
	CountPosts() (int64, error)
	FindPaginated(offset int, Limit int) ([]entities.Posts, error)
	UpdatePost(post *entities.Posts) error
	DeletePostById(post *entities.Posts) error
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

func (r *repository) UpdatePost(post *entities.Posts) error {
	return r.Db.Model(post).Updates(entities.Posts{Title: post.Title, Slug: post.Slug, Body: post.Body}).Error
}

func (r *repository) GetPosts(posts *[]entities.Posts) (*[]entities.Posts, error) {
	result := r.Db.Preload("Creator").Order("created_at DESC").Find(posts)
	err := result.Error
	if err != nil {
		return &[]entities.Posts{}, err
	}
	return posts, nil
}

func (r *repository) FindPaginated(offset int, Limit int) ([]entities.Posts, error) {
	var posts []entities.Posts
	result := r.Db.Offset(offset).Preload("Tags").Preload("Creator").Order("created_at desc").Limit(Limit).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (r *repository) GetPostsRandom(take int) (*[]entities.Posts, error) {
	posts := new([]entities.Posts)
	result := r.Db.Preload("Creator").Preload("Tags").Order("RANDOM()").Limit(take).Find(posts)
	err := result.Error
	if err != nil {
		return &[]entities.Posts{}, err
	}
	return posts, nil
}

func (r *repository) GetPost(id string, post *entities.Posts) (*entities.Posts, error) {
	err := r.Db.Preload("Creator").Where("id = ?", id).First(post).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (repo *repository) GetPostBySlug(slug string) (*entities.Posts, error) {
	post := new(entities.Posts)
	if err := repo.Db.Preload("Creator").Where("slug = ?", slug).First(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *repository) DeletePostById(post *entities.Posts) error {
	err := r.Db.Delete(post).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CountPosts() (int64, error) {
	var postCount int64
	result := r.Db.Model(&entities.Posts{}).Count(&postCount)
	return postCount, result.Error
}
