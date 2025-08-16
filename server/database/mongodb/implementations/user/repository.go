package user

import (
	"context"
	"go-vault/database/models"
	"go-vault/database/mongodb"
	"go-vault/database/repository"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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
	collection := r.mongoDB.GetDatabase().Collection(mongodb.UserCol)
	_, err := collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepositoryImpl) Update(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.UserCol)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
	return err
}

func (r *UserRepositoryImpl) DeleteByID(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.UserCol)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"deleted": true}})
	return err
}

func (r *UserRepositoryImpl) GetByID(id uuid.UUID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.UserCol)
	var user models.User
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetByUsername(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.UserCol)
	var user models.User
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
