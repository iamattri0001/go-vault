package service

import (
	customerrors "go-vault/custom_errors"
	"go-vault/database/models"
	"go-vault/utils/dbutils"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func (s *Service) CreateVault(userID uuid.UUID, request *CreateVaultRequest) (*models.Vault, error) {
	if err := validator.New().Struct(request); err != nil {
		return nil, customerrors.ErrBadRequest
	}

	if userID == uuid.Nil {
		return nil, customerrors.ErrUserNotFound
	}

	vault, err := toVaultModel(userID, request)
	if err != nil {
		return nil, err
	}

	if s.vaultRepository.ExistsByUserIdAndTitle(userID, request.Title) {
		return nil, customerrors.ErrVaultAlreadyExists
	}

	if err := s.vaultRepository.Create(vault); err != nil {
		return nil, customerrors.ErrSomethingWentWrong
	}

	return vault, nil
}

func (s *Service) ListVaults(userID uuid.UUID) ([]*models.Vault, error) {
	if userID == uuid.Nil {
		return nil, customerrors.ErrUserNotFound
	}
	vaults, err := s.vaultRepository.GetByUserID(userID)
	if err != nil {
		return nil, customerrors.ErrSomethingWentWrong
	}
	return vaults, nil
}

func (s *Service) UpdateVault(userID, vaultID uuid.UUID, request *UpdateVaultRequest) (*models.Vault, error) {
	if err := validator.New().Struct(request); err != nil {
		return nil, customerrors.ErrBadRequest
	}

	if userID == uuid.Nil {
		return nil, customerrors.ErrUserNotFound
	}

	if vaultID == uuid.Nil {
		return nil, customerrors.ErrVaultNotFound
	}

	vault, err := s.vaultRepository.GetByID(vaultID)
	if err != nil {
		log.Printf("Error fetching vault: %v", err)
		return nil, customerrors.ErrVaultNotFound
	}

	if s.vaultRepository.ExistsByUserIdAndTitle(userID, request.Title) && vault.Title != request.Title {
		return nil, customerrors.ErrVaultAlreadyExists
	}

	if vault.UserID != userID {
		return nil, customerrors.ErrUnauthorized
	}

	vault, err = toUpdatedVaultModel(vaultID, userID, request)
	if err != nil {
		return nil, err
	}

	if err := s.vaultRepository.Update(vault); err != nil {
		return nil, customerrors.ErrSomethingWentWrong
	}

	return vault, nil
}

func (s *Service) DeleteVaultAndPasswordByVaultID(userID, vaultID uuid.UUID) error {
	vault, err := s.vaultRepository.GetByID(vaultID)
	if err != nil {
		return customerrors.ErrVaultNotFound
	}

	if vault.UserID != userID {
		return customerrors.ErrUnauthorized
	}

	err = dbutils.RunParallel(
		func() error {
			return s.vaultRepository.DeleteByID(vaultID)
		},
		func() error {
			return s.passwordRepository.DeleteByVaultID(vaultID)
		},
	)

	if err != nil {
		log.Printf("Error deleting vault and passwords: %v", err)
		return customerrors.ErrSomethingWentWrong
	}
	return nil
}
