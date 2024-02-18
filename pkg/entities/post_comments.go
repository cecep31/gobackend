package entities

import (
	"gobackend/database"

	"github.com/google/uuid"
)

type PostComments struct {
	database.DefaultModel
	Text             string `json:"text"`
	PostId           string
	ParrentCommentId uint64
	CreatedBy        uuid.UUID       `json:"created_by"`
	Creator          Users           `json:"creator" gorm:"foreignKey:CreatedBy"`
	replies          []*PostComments `gorm:"foreignKey:ID"`
}
