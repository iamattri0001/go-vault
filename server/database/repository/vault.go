package repository

import (
	"go-vault/database/models"

	"github.com/google/uuid"
)

type VaultRepository interface {
	Create(vault *models.Vault) error
	Update(vault *models.Vault) error
	DeleteByID(id uuid.UUID) error
	GetByID(id uuid.UUID) (*models.Vault, error)
	GetByUserID(userID uuid.UUID) ([]*models.Vault, error)
}
