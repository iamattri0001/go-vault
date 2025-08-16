package models

import "github.com/google/uuid"

type Vault struct {
	Base        `bson:",inline"`
	UserID      uuid.UUID `json:"user_id" bson:"user_id"`
	Title       string    `json:"title" bson:"title"`
	Description *string   `json:"description,omitempty" bson:"description,omitempty"`
}
