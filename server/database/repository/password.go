package repository

import (
	"go-vault/database/models"

	"github.com/google/uuid"
)

type PasswordRepository interface {
	Create(password *models.Password) error
	Update(password *models.Password) error
	DeleteByID(id uuid.UUID) error
	DeleteByVaultID(vaultID uuid.UUID) error
	GetByID(id uuid.UUID) (*models.Password, error)
	GetByVaultID(vaultID uuid.UUID) ([]*models.Password, error)
	ExistsByVaultIDAndTitle(vaultID uuid.UUID, title string) bool
}
