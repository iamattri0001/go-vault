package models

import "github.com/google/uuid"

type Password struct {
	Base
	VaultID           uuid.UUID `json:"vault_id" bson:"vault_id"`
	Title             string    `json:"title" bson:"title"`
	Description       string    `json:"description" bson:"description"`
	Username          string    `json:"username" bson:"username"`
	EncryptedPassword string    `json:"encrypted_password" bson:"encrypted_password"`
	Website           string    `json:"website" bson:"website"`
}
