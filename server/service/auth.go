package service

import (
	customerrors "go-vault/custom_errors"
	"go-vault/database/models"
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
