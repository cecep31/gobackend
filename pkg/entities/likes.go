package entities

import (
	"time"

	"github.com/google/uuid"
)

type Likes struct {
	ID         uint64    `json:"id"`
	Post_id    uuid.UUID `json:"post_id"`
	Created_by uuid.UUID `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
}
