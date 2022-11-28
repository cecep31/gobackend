package entities

import "gobackend/database"

type User struct {
	database.DefaultModel
	Username string `json:"username" gorm:"uniqueIndex"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Role     string `json:"role"`
	Password string `json:"password" gorm:"type:text"`
	Image    string `json:"image" gorm:"type:text"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}
