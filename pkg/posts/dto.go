package posts

import (
	"gobackend/database"

	"github.com/google/uuid"
)

type Posts struct {
	database.DefaultModel
	Title     string    `json:"title" validate:"required,min=8"`
	Photo_url string    `json:"photo_url" validate:"required"`
	Body      string    `json:"body" validate:"required,min=100"`
	CreatedBy uuid.UUID `json:"created_by"`
	Slug      string    `json:"slug" validate:"required"`
}
