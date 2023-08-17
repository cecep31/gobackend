package entities

import (
	"gobackend/database"

	"github.com/google/uuid"
)

type Tasks struct {
	database.DefaultModel
	Title     string    `json:"title"`
	Body      string    `json:"body" gorm:"type:text"`
	CreatedBy uuid.UUID `json:"created_by"`
	Creator   Users     `gorm:"foreignKey:CreatedBy"`
	Order     int64     `json:"order"`
}
