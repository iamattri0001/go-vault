package password

import (
	"context"
	"go-vault/database/models"
	"go-vault/database/mongodb"
	"go-vault/database/repository"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type PasswordRepositoryImpl struct {
	mongoDB *mongodb.MongoDBConnection
}

func NewPasswordRepositoryImpl(client *mongodb.MongoDBConnection) repository.PasswordRepository {
	return &PasswordRepositoryImpl{
		mongoDB: client,
	}
}

func (r *PasswordRepositoryImpl) Create(password *models.Password) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.PasswordCol)
	_, err := collection.InsertOne(ctx, password)
	return err
}

func (r *PasswordRepositoryImpl) Update(password *models.Password) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.PasswordCol)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": password.ID}, bson.M{"$set": password})
	return err
}

func (r *PasswordRepositoryImpl) DeleteByID(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.PasswordCol)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"deleted_at": time.Now()}})
	return err
}

func (r *PasswordRepositoryImpl) GetByID(id uuid.UUID) (*models.Password, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.PasswordCol)
	var password models.Password
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&password)
	if err != nil {
		return nil, err
	}
	return &password, nil
}

func (r *PasswordRepositoryImpl) GetByCollectionID(collectionID uuid.UUID) ([]*models.Password, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.PasswordCol)
	var passwords []*models.Password
	cursor, err := collection.Find(ctx, bson.M{"collection_id": collectionID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var password models.Password
		if err := cursor.Decode(&password); err != nil {
			return nil, err
		}
		passwords = append(passwords, &password)
	}
	return passwords, nil
}
