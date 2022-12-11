package entities

import "gobackend/database"

type Taskgorup struct {
	database.DefaultModel
	Name       string `json:"name"`
	Created_by int64  `json:"created_by"`
	CreatedBy  User   `gorm:"foreignKey:created_by"`
	Task       []Task `gorm:"foreignKey:GroupID"`
	Order      int64  `json:"order"`
}
