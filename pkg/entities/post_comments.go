package entities

import (
	"gobackend/database"
)

type PostComments struct {
	database.SecondDefaultModel
	Text             string `json:"text"`
	PostId           string
	ParrentCommentId uint64
	replies          []*PostComments `gorm:"foreignKey:id"`
}
