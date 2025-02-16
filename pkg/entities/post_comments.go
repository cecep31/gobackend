package entities

import (
	"gobackend/database/template"

	"github.com/google/uuid"
)

type PostComments struct {
	template.DefaultModel
	Text             string `json:"text"`
	PostId           string
	ParrentCommentId uint64
	CreatedBy        uuid.UUID       `json:"created_by"`
	Creator          User            `json:"creator" gorm:"foreignKey:CreatedBy"`
	replies          []*PostComments `gorm:"foreignKey:ID"`
}
