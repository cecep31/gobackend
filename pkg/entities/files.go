package entities

import (
	"gobackend/database/template"
	"time"

	"github.com/google/uuid"
)

type Files struct {
	template.DefaultModel
	Name       string     `json:"name"`
	Path       string     `json:"path"`
	Size       int64      `json:"size"`
	Type       string     `json:"type"`
	Created_by uuid.UUID  `json:"created_by"`
	Deleted_at *time.Time `json:"deleted_at"`
	Uploader   User       `gorm:"foreignKey:Created_by"`
}
