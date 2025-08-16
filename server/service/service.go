package service

import "go-vault/database/repository"

type Service struct {
	userRepository     repository.UserRepository
	passwordRepository repository.PasswordRepository
	vaultRepository    repository.VaultRepository
}

func NewService(userRepo repository.UserRepository, vaultRepo repository.VaultRepository, passwordRepo repository.PasswordRepository) *Service {
	return &Service{
		userRepository:     userRepo,
		passwordRepository: passwordRepo,
		vaultRepository:    vaultRepo,
	}
}
