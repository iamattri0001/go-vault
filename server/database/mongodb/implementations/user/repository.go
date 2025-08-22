package user

import (
	"context"
	"go-vault/database/models"
	"go-vault/database/mongodb"
	"go-vault/database/repository"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositoryImpl struct {
	mongoDB *mongodb.MongoDBConnection
}

func NewUserRepositoryImpl(client *mongodb.MongoDBConnection) repository.UserRepository {
	return &UserRepositoryImpl{
		mongoDB: client,
	}
}

func (r *UserRepositoryImpl) Create(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.UserCollection)
	_, err := collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepositoryImpl) Update(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.UserCollection)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
	return err
}

func (r *UserRepositoryImpl) DeleteByID(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.UserCollection)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"deleted_at": time.Now()}})
	return err
}

func (r *UserRepositoryImpl) GetByID(id uuid.UUID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.UserCollection)
	var user models.User
	err := collection.FindOne(ctx, bson.M{"_id": id, "deleted_at": nil}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetByUsername(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.UserCollection)
	var user models.User
	err := collection.FindOne(ctx, bson.M{"username": username, "deleted_at": nil}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) ExistsByUsername(username string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.UserCollection)
	var user models.User
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return err == nil && user.ID != uuid.Nil
}

func (r *UserRepositoryImpl) GetSaltsByUsername(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.UserCollection)
	var user models.User
	err := collection.FindOne(ctx, bson.M{"username": username, "deleted_at": nil}, options.FindOne().SetProjection(bson.M{"auth_salt": 1, "encryption_salt": 1})).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
