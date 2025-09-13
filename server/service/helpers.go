package service

import (
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

	passwordHash, err := hash.GenerateHash(request.Password)
	if err != nil {
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
		AuthSalt:       request.AuthSalt,
		EncryptionSalt: request.EncryptionSalt,
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

func toUpdatedVaultModel(userID uuid.UUID, request *UpdateVaultRequest) (*models.Vault, error) {
	if len(request.Title) < 3 || len(request.Title) > 30 {
		return nil, customerrors.ErrInvalidTitleFormat
	}

	if request.Description != nil && len(*request.Description) > 100 {
		return nil, customerrors.ErrInvalidDescriptionFormat
	}

	return &models.Vault{
		Base: models.Base{
			ID:        request.ID,
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

func toUpdatedPasswordModel(userID uuid.UUID, vaultID uuid.UUID, request *UpdatePasswordRequest) (*models.Password, error) {
	if len(request.Title) < 3 || len(request.Title) > 30 {
		return nil, customerrors.ErrInvalidTitleFormat
	}

	if request.Description != nil && len(*request.Description) > 100 {
		return nil, customerrors.ErrInvalidDescriptionFormat
	}

	return &models.Password{
		Base: models.Base{
			ID:        request.ID,
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
