package entities

import (
	"gobackend/database"

	"github.com/google/uuid"
)

type Globalchat struct {
	database.DefaultModel
	UserID uuid.UUID
	User   Users
	Msg    string
}
