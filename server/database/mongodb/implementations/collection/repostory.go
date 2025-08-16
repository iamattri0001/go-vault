package collection

import (
	"context"
	"go-vault/database/models"
	"go-vault/database/mongodb"
	"go-vault/database/repository"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type CollectionRepositoryImpl struct {
	mongoDB *mongodb.MongoDBConnection
}

func NewCollectionRepositoryImpl(client *mongodb.MongoDBConnection) repository.CollectionRepository {
	return &CollectionRepositoryImpl{
		mongoDB: client,
	}
}

func (r *CollectionRepositoryImpl) Create(collection *models.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	mongoCollection := r.mongoDB.GetDatabase().Collection(mongodb.CollectionCol)
	_, err := mongoCollection.InsertOne(ctx, collection)
	return err
}

func (r *CollectionRepositoryImpl) Update(collection *models.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	mongoCollection := r.mongoDB.GetDatabase().Collection(mongodb.CollectionCol)
	_, err := mongoCollection.UpdateOne(ctx, bson.M{"_id": collection.ID}, bson.M{"$set": collection})
	return err
}

func (r *CollectionRepositoryImpl) DeleteByID(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	mongoCollection := r.mongoDB.GetDatabase().Collection(mongodb.CollectionCol)
	_, err := mongoCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"deleted": true}})
	return err
}

func (r *CollectionRepositoryImpl) GetByID(id uuid.UUID) (*models.Collection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	mongoCollection := r.mongoDB.GetDatabase().Collection(mongodb.CollectionCol)
	var collection models.Collection
	err := mongoCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&collection)
	if err != nil {
		return nil, err
	}
	return &collection, nil
}

func (r *CollectionRepositoryImpl) GetByUserID(userID uuid.UUID) ([]*models.Collection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	mongoCollection := r.mongoDB.GetDatabase().Collection(mongodb.CollectionCol)
	var collections []*models.Collection
	cursor, err := mongoCollection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var collection models.Collection
		if err := cursor.Decode(&collection); err != nil {
			return nil, err
		}
		collections = append(collections, &collection)
	}
	return collections, nil
}
