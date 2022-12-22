package entities

import "gobackend/database"

type Post struct {
	database.DefaultModel
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Created_by int64  `json:"created_by"`
	CreatedBy  User   `gorm:"foreignKey:Created_by"`
}
