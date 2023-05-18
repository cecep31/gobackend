package entities

import "gobackend/database"

type Users struct {
	database.DefaultModel
	Username     string `json:"username" gorm:"uniqueIndex"`
	Name         string `json:"name" gorm:"default:pilput"`
	Email        string `json:"email" gorm:"uniqueIndex"`
	Password     string `json:"password" gorm:"type:text"`
	Image        string `json:"image" gorm:"type:text"`
	Issuperadmin bool   `json:"issuperadmin" gorm:"default:false"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}
