package service

import (
	"crypto/rand"
	customerrors "go-vault/custom_errors"
	"go-vault/database/models"
	"go-vault/pkg/hash"
	"time"

	"github.com/google/uuid"
)

func toUserModel(request *CreateUserRequest) (*models.User, error) {
	if len(request.Username) < 3 || len(request.Username) > 15 {
		return nil, customerrors.ErrInvalidUsernameFormat
	}

	if len(request.Password) < 8 || len(request.Password) > 30 {
		return nil, customerrors.ErrInvalidPasswordFormat
	}

	passwordHash, err := hash.GenerateHash(request.Password)
	if err != nil {
		return nil, customerrors.ErrSomethingWentWrong
	}

	authSalt := make([]byte, 16)
	if _, err := rand.Read(authSalt); err != nil {
		return nil, customerrors.ErrSomethingWentWrong
	}

	encryptionSalt := make([]byte, 32)
	if _, err := rand.Read(encryptionSalt); err != nil {
		return nil, customerrors.ErrSomethingWentWrong
	}

	return &models.User{
		Base: models.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username:       request.Username,
		PasswordHash:   passwordHash,
		AuthSalt:       string(authSalt),
		EncryptionSalt: string(encryptionSalt),
	}, nil
}

func toVaultModel(userID uuid.UUID, request *CreateVaultRequest) (*models.Vault, error) {
	if len(request.Title) < 3 || len(request.Title) > 30 {
		return nil, customerrors.ErrInvalidTitleFormat
	}

	if request.Description != nil && len(*request.Description) > 100 {
		return nil, customerrors.ErrInvalidDescriptionFormat
	}

	return &models.Vault{
		Base: models.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UserID:      userID,
		Title:       request.Title,
		Description: request.Description,
	}, nil
}

func toUpdatedVaultModel(vaultID uuid.UUID, userID uuid.UUID, request *UpdateVaultRequest) (*models.Vault, error) {
	if len(request.Title) < 3 || len(request.Title) > 30 {
		return nil, customerrors.ErrInvalidTitleFormat
	}

	if request.Description != nil && len(*request.Description) > 100 {
		return nil, customerrors.ErrInvalidDescriptionFormat
	}

	return &models.Vault{
		Base: models.Base{
			ID:        vaultID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Title:       request.Title,
		Description: request.Description,
		UserID:      userID,
	}, nil
}

func toPasswordModel(userID uuid.UUID, vaultID uuid.UUID, request *CreatePasswordRequest) (*models.Password, error) {
	if len(request.Title) < 3 || len(request.Title) > 30 {
		return nil, customerrors.ErrInvalidTitleFormat
	}

	if request.Description != nil && len(*request.Description) > 100 {
		return nil, customerrors.ErrInvalidDescriptionFormat
	}

	return &models.Password{
		Base: models.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UserID:            userID,
		VaultID:           vaultID,
		Title:             request.Title,
		Description:       request.Description,
		Username:          request.Username,
		EncryptedPassword: request.Password,
		Website:           request.Website,
	}, nil
}
