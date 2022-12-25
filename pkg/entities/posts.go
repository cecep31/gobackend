package entities

import "gobackend/database"

type Posts struct {
	database.DefaultModel
	Title      string      `json:"title"`
	Desc       string      `json:"desc"`
	Created_by int64       `json:"created_by"`
	CreatedBy  Users       `gorm:"foreignKey:Created_by"`
	Posttags   []*Posttags `gorm:"many2many:posts_posttags"`
}
