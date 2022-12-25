package entities

import "gobackend/database"

type Posttags struct {
	database.DefaultModel
	Name  string   `json:"name"`
	Posts []*Posts `gorm:"many2many:posts_posttags"`
}
