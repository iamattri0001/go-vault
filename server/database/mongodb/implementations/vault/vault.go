package vault

import (
	"context"
	"go-vault/database/models"
	"go-vault/database/mongodb"
	"go-vault/database/repository"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type VaultRepositoryImpl struct {
	mongoDB *mongodb.MongoDBConnection
}

func NewVaultRepositoryImpl(client *mongodb.MongoDBConnection) repository.VaultRepository {
	return &VaultRepositoryImpl{
		mongoDB: client,
	}
}

func (r *VaultRepositoryImpl) Create(vault *models.Vault) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.VaultCollection)
	_, err := collection.InsertOne(ctx, vault)
	return err
}

func (r *VaultRepositoryImpl) Update(vault *models.Vault) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.VaultCollection)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": vault.ID}, bson.M{"$set": vault})
	return err
}

func (r *VaultRepositoryImpl) DeleteByID(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.VaultCollection)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"deleted_at": time.Now()}})
	return err
}

func (r *VaultRepositoryImpl) GetByID(id uuid.UUID) (*models.Vault, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.VaultCollection)
	var vault models.Vault
	err := collection.FindOne(ctx, bson.M{"_id": id, "deleted_at": nil}).Decode(&vault)
	if err != nil {
		return nil, err
	}
	return &vault, nil
}

func (r *VaultRepositoryImpl) GetByUserID(userID uuid.UUID) ([]*models.Vault, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	mongoCollection := r.mongoDB.GetDatabase().Collection(mongodb.VaultCollection)
	vaults := make([]*models.Vault, 0)
	cursor, err := mongoCollection.Find(ctx, bson.M{"user_id": userID, "deleted_at": nil})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var vault models.Vault
		if err := cursor.Decode(&vault); err != nil {
			return nil, err
		}
		vaults = append(vaults, &vault)
	}
	return vaults, nil
}

func (r *VaultRepositoryImpl) ExistsByUserIdAndTitle(userID uuid.UUID, title string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.VaultCollection)
	var vault models.Vault
	err := collection.FindOne(ctx, bson.M{"user_id": userID, "title": title, "deleted_at": nil}).Decode(&vault)
	return err == nil && vault.ID != uuid.Nil
}
