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

	user, err := toUserModel(request)
	if err != nil {
		return nil, err
	}

	if s.userRepository.ExistsByUsername(request.Username) {
		return nil, customerrors.ErrUsernameTaken
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

	return user, nil
}
