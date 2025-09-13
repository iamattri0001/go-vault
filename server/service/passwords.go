package service

import (
	customerrors "go-vault/custom_errors"
	"go-vault/database/models"
	"log"

	"github.com/google/uuid"
)

func (s *Service) GetVaultPasswords(userID uuid.UUID, vaultID uuid.UUID) ([]*models.Password, error) {
	vault, err := s.vaultRepository.GetByID(vaultID)
	if err != nil {
		log.Printf("Error fetching vault by ID %s: %v", vaultID, err)
		return nil, customerrors.ErrSomethingWentWrong
	}

	if vault.UserID != userID {
		return nil, customerrors.ErrUnauthorized
	}

	passwords, err := s.passwordRepository.GetByVaultID(vaultID)
	if err != nil {
		log.Printf("Error fetching passwords for vault ID %s: %v", vaultID, err)
		return nil, customerrors.ErrSomethingWentWrong
	}

	return passwords, nil
}

func (s *Service) CreatePasswordWithVaultID(userID uuid.UUID, request *CreatePasswordRequest) (*models.Password, error) {
	if userID == uuid.Nil {
		return nil, customerrors.ErrUserNotFound
	}

	vaultID := request.VaultID
	if vaultID == uuid.Nil {
		return nil, customerrors.ErrVaultNotFound
	}

	vault, err := s.vaultRepository.GetByID(vaultID)
	if err != nil {
		return nil, customerrors.ErrSomethingWentWrong
	}
	if vault.UserID != userID {
		return nil, customerrors.ErrUnauthorized
	}

	password, err := toPasswordModel(userID, vaultID, request)
	if err != nil {
		return nil, err
	}

	if s.passwordRepository.ExistsByVaultIDAndTitle(vaultID, password.Title) {
		return nil, customerrors.ErrPasswordAlreadyExists
	}

	if err := s.passwordRepository.Create(password); err != nil {
		log.Printf("Error creating password: %v", err)
		return nil, customerrors.ErrSomethingWentWrong
	}

	return password, nil
}

func (s *Service) UpdatePassword(userID uuid.UUID, request *UpdatePasswordRequest) (*models.Password, error) {
	if userID == uuid.Nil {
		return nil, customerrors.ErrUserNotFound
	}
	vaultID := request.VaultID
	if vaultID == uuid.Nil {
		return nil, customerrors.ErrVaultNotFound
	}

	vault, err := s.vaultRepository.GetByID(vaultID)
	if err != nil {
		return nil, customerrors.ErrSomethingWentWrong
	}

	if vault.UserID != userID {
		return nil, customerrors.ErrUnauthorized
	}
	password, err := s.passwordRepository.GetByID(request.ID)
	if err != nil {
		return nil, customerrors.ErrSomethingWentWrong
	}

	if password.UserID != userID {
		return nil, customerrors.ErrUnauthorized
	}

	updatedPassword, err := toUpdatedPasswordModel(userID, vaultID, request)
	if err != nil {
		return nil, err
	}

	if err := s.passwordRepository.Update(updatedPassword); err != nil {
		log.Printf("Error updating password: %v", err)
		return nil, customerrors.ErrSomethingWentWrong
	}

	return updatedPassword, nil
}
