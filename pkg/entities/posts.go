package entities

import (
	"gobackend/database"

	"github.com/google/uuid"
)

type Posts struct {
	database.DefaultModel
	Title     string      `json:"title"`
	Body      string      `json:"body" gorm:"type=text"`
	CreatedBy uuid.UUID   `json:"created_by" gorm:"type=uuid"`
	User      Users       `gorm:"foreignKey:CreatedBy"`
	Posttags  []*Posttags `gorm:"many2many:posts_posttags"`
}
