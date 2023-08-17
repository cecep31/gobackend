package entities

import (
	"gobackend/database"

	"github.com/google/uuid"
)

type Posts struct {
	database.DefaultModel
	Title        string         `json:"title"`
	Slug         string         `json:"slug" gorm:"unique"`
	Body         string         `json:"body" gorm:"type=text"`
	CreatedBy    uuid.UUID      `json:"created_by"`
	Creator      Users          `gorm:"foreignKey:CreatedBy"`
	PostComments []PostComments `gorm:"foreignKey:PostId"`
}
