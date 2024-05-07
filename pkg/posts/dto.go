package posts

import (
	"gobackend/database/template"

	"github.com/google/uuid"
)

type PostCreate struct {
	template.DefaultModel
	Title     string    `json:"title" validate:"required,min=8"`
	Photo_url string    `json:"photo_url" validate:"required"`
	Body      string    `json:"body" validate:"required,min=50"`
	CreatedBy uuid.UUID `json:"created_by"`
	Slug      string    `json:"slug" validate:"required"`
}
type PostUpdate struct {
	template.DefaultModel
	Title     string    `json:"title" validate:"required,min=8"`
	Photo_url string    `json:"photo_url" validate:"required"`
	Body      string    `json:"body" validate:"required,min=50"`
	CreatedBy uuid.UUID `json:"created_by"`
	Slug      string    `json:"slug" validate:"required"`
}
