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
func (r *repository) GetPosts() (*[]entities.Posts, error) {
	var posts []entities.Posts
	result := r.Db.Preload("Creator").Order("created_at DESC").Find(&posts)
	err := result.Error
	if err != nil {
		return &[]entities.Posts{}, err
	}
	return &posts, nil
}

func (r *repository) GetPostsRandom(take int) (*[]entities.Posts, error) {
	posts := new([]entities.Posts)
	result := r.Db.Preload("Creator").Order("RANDOM()").Take(take).Find(posts)
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
