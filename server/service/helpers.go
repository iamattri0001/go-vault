package service

import (
	customerrors "go-vault/custom_errors"
	"go-vault/database/models"
	"go-vault/pkg/hash"
	"time"

	"github.com/google/uuid"
)

func toUserModel(request *CreateUserRequest) (*models.User, error) {
	passwordHash, err := hash.GenerateHash(request.Password)
	if err != nil {
		return nil, customerrors.ErrSomethingWentWrong
	}

	if len(request.Username) < 3 || len(request.Username) > 15 {
		return nil, customerrors.ErrInvalidUsernameFormat
	}

	if len(request.Password) < 8 || len(request.Password) > 30 {
		return nil, customerrors.ErrInvalidPasswordFormat
	}

	return &models.User{
		Base: models.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username:           request.Username,
		MasterPasswordHash: passwordHash,
	}, nil
}
