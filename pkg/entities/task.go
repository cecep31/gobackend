package entities

import (
	"gobackend/database"

	"github.com/google/uuid"
)

type Tasks struct {
	database.DefaultModel
	Title      string `json:"title"`
	Desc       string `json:"desc" gorm:"type:text"`
	GroupID    int64  `json:"group_id"`
	Group      Taskgorups
	Created_by uuid.UUID `json:"created_by"`
	CreatedBy  Users     `gorm:"foreignKey:Created_by"`
	Order      int64     `json:"order"`
}
