package models

import "github.com/google/uuid"

type Password struct {
	Base              `bson:",inline"`
	VaultID           uuid.UUID `json:"vault_id" bson:"vault_id"`
	UserID            uuid.UUID `json:"user_id" bson:"user_id"`
	Title             string    `json:"title" bson:"title"`
	Description       *string   `json:"description,omitempty" bson:"description,omitempty"`
	Username          string    `json:"username" bson:"username"`
	EncryptedPassword string    `json:"encrypted_password" bson:"encrypted_password"`
	Website           string    `json:"website" bson:"website"`
}
