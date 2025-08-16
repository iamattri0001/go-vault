package service

import (
	customerrors "go-vault/custom_errors"
	"go-vault/database/models"
	"go-vault/pkg/hash"
	"log"

	"github.com/go-playground/validator/v10"
)

func (s *Service) CreateUser(request *CreateUserRequest) (*models.User, error) {
	if err := validator.New().Struct(request); err != nil {
		return nil, customerrors.ErrBadRequest
	}

	if s.userRepository.ExistsByUsername(request.Username) {
		return nil, customerrors.ErrUsernameTaken
	}

	user, err := toUserModel(request)
	if err != nil {
		return nil, err
	}

	err = s.userRepository.Create(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, customerrors.ErrSomethingWentWrong
	}
	return user, nil
}

func (s *Service) LoginUser(request *LoginUserRequest) (*models.User, error) {
	if err := validator.New().Struct(request); err != nil {
		return nil, customerrors.ErrBadRequest
	}

	user, err := s.userRepository.GetByUsername(request.Username)
	if err != nil {
		return nil, customerrors.ErrUserNotFound
	}

	if err := hash.CheckHash(request.Password, user.MasterPasswordHash); err != nil {
		log.Printf("Login failed for user %s: %v", request.Username, err)
		return nil, customerrors.ErrInvalidCredentials
	}

	return user, nil
}
