package entities

import (
	"gobackend/database"
)

type Books struct {
	database.DefaultModel
	Title      string `json:"title"`
	Desc       string `json:"desc" gorm:"type:text"`
	Author     string `json:"author"`
	Rating     int    `json:"rating"`
	Created_by int64  `json:"create_at"`
	CreatedBy  Users  `gorm:"foreignKey:Created_by"`
	Price      int32  `json:"price" gorm:"default:0"`
	Image      string `json:"image" gorm:"type:text"`
}
