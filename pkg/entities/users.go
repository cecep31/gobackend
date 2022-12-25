package entities

import "gobackend/database"

type Users struct {
	database.DefaultModel
	Username     string `json:"username" gorm:"uniqueIndex"`
	Email        string `json:"email" gorm:"uniqueIndex"`
	Role         string `json:"role"`
	Password     string `json:"password" gorm:"type:text"`
	Image        string `json:"image" gorm:"type:text"`
	Issuperadmin bool   `json:"issuperadmin" gorm:"default:false"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}
