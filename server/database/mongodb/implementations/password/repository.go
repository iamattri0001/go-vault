package password

import (
	"context"
	"go-vault/database/models"
	"go-vault/database/mongodb"
	"go-vault/database/repository"

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
	collection := r.mongoDB.GetDatabase().Collection(mongodb.PasswordCollection)
	_, err := collection.InsertOne(ctx, password)
	return err
}

func (r *PasswordRepositoryImpl) Update(password *models.Password) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.PasswordCollection)
	_, err := collection.UpdateOne(ctx, bson.M{"_id": password.ID}, bson.M{"$set": password})
	return err
}

func (r *PasswordRepositoryImpl) DeleteByID(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.PasswordCollection)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *PasswordRepositoryImpl) GetByID(id uuid.UUID) (*models.Password, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.PasswordCollection)
	var password models.Password
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&password)
	if err != nil {
		return nil, err
	}
	return &password, nil
}

func (r *PasswordRepositoryImpl) GetByVaultID(vaultID uuid.UUID) ([]*models.Password, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.PasswordCollection)
	passwords := make([]*models.Password, 0)
	cursor, err := collection.Find(ctx, bson.M{"vault_id": vaultID})
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

func (r *PasswordRepositoryImpl) DeleteByVaultID(vaultID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.PasswordCollection)
	_, err := collection.DeleteMany(ctx, bson.M{"vault_id": vaultID})
	return err
}

func (r *PasswordRepositoryImpl) ExistsByVaultIDAndTitle(vaultID uuid.UUID, title string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), r.mongoDB.QueryTimeout)
	defer cancel()
	collection := r.mongoDB.GetDatabase().Collection(mongodb.PasswordCollection)
	count, _ := collection.CountDocuments(ctx, bson.M{"vault_id": vaultID, "title": title, "deleted_at": nil})
	return count > 0
}
