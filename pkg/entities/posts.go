package entities

import (
	"gobackend/database"

	"github.com/google/uuid"
)

type Posts struct {
	database.DefaultModel
	Title     string      `json:"title"`
	Desc      string      `json:"desc"`
	CreatedBy uuid.UUID   `json:"created_by"`
	User      Users       `gorm:"foreignKey:CreatedBy"`
	Posttags  []*Posttags `gorm:"many2many:posts_posttags"`
}
