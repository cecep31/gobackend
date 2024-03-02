package entities

import (
	"gobackend/database/template"

	"github.com/google/uuid"
)

type Tasks struct {
	template.DefaultModel
	Title     string    `json:"title"`
	Body      string    `json:"body" gorm:"type:text"`
	CreatedBy uuid.UUID `json:"created_by"`
	Creator   Users     `gorm:"foreignKey:CreatedBy"`
	Order     int64     `json:"order"`
}
