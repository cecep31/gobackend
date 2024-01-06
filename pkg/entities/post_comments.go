package entities

import (
	"gobackend/database"
)

type PostComments struct {
	database.DefaultModel
	Text             string `json:"text"`
	PostId           string
	ParrentCommentId uint64
	replies          []*PostComments `gorm:"foreignKey:ID"`
}
