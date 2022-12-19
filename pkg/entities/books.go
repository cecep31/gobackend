package entities

import (
	"gobackend/database"
)

type Book struct {
	database.DefaultModel
	Title  string `json:"title"`
	Desc   string `json:"desc" gorm:"type:text"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
	Price  int32  `json:"price" gorm:"default:0"`
	Image  string `json:"image" gorm:"type:text"`
}
