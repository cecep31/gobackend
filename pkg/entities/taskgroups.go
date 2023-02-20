package entities

import (
	"gobackend/database"

	"github.com/google/uuid"
)

type Taskgorups struct {
	database.DefaultModel
	Name       string    `json:"name"`
	Created_by uuid.UUID `json:"created_by"`
	CreatedBy  Users     `gorm:"foreignKey:created_by"`
	Task       []Tasks   `gorm:"foreignKey:GroupID"`
	Order      int64     `json:"order"`
}
