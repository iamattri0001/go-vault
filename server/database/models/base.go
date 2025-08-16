package models

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID  `json:"id" bson:"_id"`
	CreatedAt time.Time  `json:"-" bson:"created_at"`
	UpdatedAt time.Time  `json:"-" bson:"updated_at"`
	DeletedAt *time.Time `json:"-" bson:"deleted_at,omitempty"`
}
