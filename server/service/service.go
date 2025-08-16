package service

import "go-vault/database/repository"

type Service struct {
	UserRepository     repository.UserRepository
	PasswordRepository repository.PasswordRepository
	VaultRepository    repository.VaultRepository
}

func NewService(userRepo repository.UserRepository, vaultRepo repository.VaultRepository, passwordRepo repository.PasswordRepository) *Service {
	return &Service{
		UserRepository:     userRepo,
		PasswordRepository: passwordRepo,
		VaultRepository:    vaultRepo,
	}
}
