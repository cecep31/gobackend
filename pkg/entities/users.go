package entities

import "gobackend/database/template"

type Users struct {
	template.DefaultModel
	Username     string `json:"username" gorm:"uniqueIndex;"`
	FirstName    string `json:"first_name" gorm:"default:pilput"`
	LastName     string `json:"last_name" gorm:"default:admin"`
	Email        string `json:"email" gorm:"uniqueIndex; not null"`
	Password     string `json:"_" gorm:"type:text"`
	Image        string `json:"image" gorm:"type:text"`
	Issuperadmin bool   `json:"issuperadmin" gorm:"default:false"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}
