package models

import "github.com/google/uuid"

type Password struct {
	Base
	CollectionID      uuid.UUID `json:"collection_id" bson:"collection_id"`
	Title             string    `json:"title" bson:"title"`
	Description       string    `json:"description" bson:"description"`
	Username          string    `json:"username" bson:"username"`
	EncryptedPassword string    `json:"encrypted_password" bson:"encrypted_password"`
	Website           string    `json:"website" bson:"website"`
}
