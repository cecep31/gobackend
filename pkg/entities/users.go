package entities

import "gobackend/database"

type Users struct {
	database.DefaultModel
	FirstName    string `json:"first_name" gorm:"default:pilput"`
	LastName     string `json:"last_name" gorm:"default:admin"`
	Email        string `json:"email" gorm:"uniqueIndex"`
	Password     string `json:"password" gorm:"type:text"`
	Image        string `json:"image" gorm:"type:text"`
	Issuperadmin bool   `json:"issuperadmin" gorm:"default:false"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}
