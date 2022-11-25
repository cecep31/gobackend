package entities

import "gobackend/database"

type Items struct {
	database.DefaultModel
	Name  string `json:"name"`
	Desc  string `json:"desc" gorm:"type:text"`
	Image string `json:"image" gorm:"type:text"`
}
