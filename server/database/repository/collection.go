package repository

import (
	"go-vault/database/models"

	"github.com/google/uuid"
)

type CollectionRepository interface {
	Create(collection *models.Collection) error
	Update(collection *models.Collection) error
	DeleteByID(id uuid.UUID) error
	GetByID(id uuid.UUID) (*models.Collection, error)
	GetByUserID(userID uuid.UUID) ([]*models.Collection, error)
}
