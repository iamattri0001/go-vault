package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ensureIndexes(db *mongo.Client, dbName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userCol := db.Database(dbName).Collection(UserCollection)
	vaultCol := db.Database(dbName).Collection(VaultCollection)
	passwordCol := db.Database(dbName).Collection(PasswordCollection)

	// User: unique username
	_, err := userCol.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		log.Fatal("Could not create user index:", err)
	}

	// Collection: index on user_id
	_, err = vaultCol.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "user_id", Value: 1}},
	})
	if err != nil {
		log.Fatal("Could not create collection index:", err)
	}

	// Password: index on collection_id
	_, err = passwordCol.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "collection_id", Value: 1}},
	})
	if err != nil {
		log.Fatal("Could not create password index:", err)
	}
}
