package entities

import (
	"gobackend/database/template"

	"github.com/google/uuid"
)

type Posts struct {
	template.DefaultModel
	Title        string         `json:"title"`
	Photo_url    string         `json:"photo_url" gorm:"type:text"`
	Slug         string         `json:"slug" gorm:"unique"`
	Body         string         `json:"body" gorm:"type=text"`
	CreatedBy    uuid.UUID      `json:"created_by"`
	Creator      Users          `json:"creator" gorm:"foreignKey:CreatedBy"`
	PostComments []PostComments `json:"comments" gorm:"foreignKey:PostId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Tags         []Tags         `json:"tags" gorm:"many2many:posts_to_tags;"`
}
