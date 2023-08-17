package entities

import (
	"gobackend/database"
)

type PostComments struct {
	database.DefaultModel
	Text             string `json:"text"`
	PostId           string
	ParrentCommentId string
	ChildComment     []*PostComments `gorm:"foreignKey:id"`
}
