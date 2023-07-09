package tasks

import (
	"gobackend/pkg/entities"

	"gorm.io/gorm"
)

type Repository interface {
	GetTasks() (*[]entities.Tasks, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetTasks() (*[]entities.Tasks, error) {
	var tasks []entities.Tasks
	err := r.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return &tasks, nil
}
