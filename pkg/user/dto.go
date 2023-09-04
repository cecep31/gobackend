package user

import "gobackend/database"

type Users struct {
	database.DefaultModel
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=8"`
	Image        string `json:"image" validate:"required"`
	Issuperadmin bool   `json:"issuperadmin"`
}
