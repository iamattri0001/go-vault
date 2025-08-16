package injector

import (
	"go-vault/database/mongodb"
	"go-vault/database/mongodb/implementations/password"
	"go-vault/database/mongodb/implementations/user"
	"go-vault/database/mongodb/implementations/vault"
	"go-vault/service"
)

func WithMongoDB(mongoDB *mongodb.MongoDBConnection) func(di *DependencyInjector) {
	return func(di *DependencyInjector) {
		userRepo := user.NewUserRepositoryImpl(mongoDB)
		passwordRepo := password.NewPasswordRepositoryImpl(mongoDB)
		vaultRepo := vault.NewVaultRepositoryImpl(mongoDB)

		service := service.NewService(userRepo, vaultRepo, passwordRepo)

		di.Service = service
	}
}
