package service

import "go-vault/database/repository"

type Service struct {
	UserRepository       repository.UserRepository
	PasswordRepository   repository.PasswordRepository
	CollectionRepository repository.CollectionRepository
}

func NewService(userRepo repository.UserRepository, collectionRepo repository.CollectionRepository, passwordRepo repository.PasswordRepository) *Service {
	return &Service{
		UserRepository:       userRepo,
		PasswordRepository:   passwordRepo,
		CollectionRepository: collectionRepo,
	}
}
