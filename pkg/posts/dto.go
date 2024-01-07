package posts

import (
	"gobackend/database"

	"github.com/google/uuid"
)

type Posts struct {
	database.DefaultModel
	Title     string    `json:"title" validate:"required"`
	Image_url string    `json:"image" validate:"required"`
	Body      string    `json:"body" validate:"required,min=100"`
	CreatedBy uuid.UUID `json:"created_by"`
	Slug      string    `json:"slug" validate:"required,min=8"`
}
