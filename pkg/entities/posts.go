package entities

import (
	"gobackend/database"

	"github.com/google/uuid"
)

type Posts struct {
	database.DefaultModel
	Title      string      `json:"title"`
	Desc       string      `json:"desc"`
	Created_by uuid.UUID   `json:"created_by"`
	CreatedBy  Users       `gorm:"foreignKey:Created_by"`
	Posttags   []*Posttags `gorm:"many2many:posts_posttags"`
}
