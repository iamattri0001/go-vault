package service

import (
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Username       string `json:"username" validate:"required"`
	Password       string `json:"password" validate:"required"`
	AuthSalt       string `json:"auth_salt" validate:"required"`
	EncryptionSalt string `json:"encryption_salt" validate:"required"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateVaultRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description *string `json:"description"`
}

type UpdateVaultRequest struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Description *string   `json:"description"`
}

type CreatePasswordRequest struct {
	Title       string    `json:"title" validate:"required"`
	VaultID     uuid.UUID `json:"vault_id" validate:"required"`
	Description *string   `json:"description"`
	Username    string    `json:"username" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	Website     string    `json:"website" validate:"required"`
}

type UpdatePasswordRequest struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	VaultID     uuid.UUID `json:"vault_id" validate:"required"`
	Description *string   `json:"description"`
	Username    string    `json:"username" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	Website     string    `json:"website" validate:"required"`
}

type Salts struct {
	AuthSalt       string `json:"auth_salt"`
	EncryptionSalt string `json:"encryption_salt"`
}
