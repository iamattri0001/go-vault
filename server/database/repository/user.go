package repository

import (
	"go-vault/database/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *models.User) error
	Update(user *models.User) error
	DeleteByID(id uuid.UUID) error
	GetByID(id uuid.UUID) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
}
