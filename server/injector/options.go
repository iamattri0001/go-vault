package injector

import (
	"go-vault/database/mongoDB/implementations/collection"
	"go-vault/database/mongodb"
	"go-vault/database/mongodb/implementations/password"
	"go-vault/database/mongodb/implementations/user"
	"go-vault/service"
)

func WithMongoDB(mongoDB *mongodb.MongoDBConnection) func(di *DependencyInjector) {
	return func(di *DependencyInjector) {
		userRepo := user.NewUserRepositoryImpl(mongoDB)
		passwordRepo := password.NewPasswordRepositoryImpl(mongoDB)
		collectionRepo := collection.NewCollectionRepositoryImpl(mongoDB)

		service := service.NewService(userRepo, collectionRepo, passwordRepo)

		di.Service = service
	}
}
