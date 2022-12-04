package entities

import "gobackend/database"

type Task struct {
	database.DefaultModel
	Title      string `json:"title"`
	Desc       string `json:"desc" gorm:"type:text"`
	GroupID    int64  `json:"group_id"`
	Group      Taskgorup
	Created_by int64 `json:"created_by"`
	CreatedBy  User  `gorm:"foreignKey:Created_by"`
	Order      int64 `json:"order"`
}
