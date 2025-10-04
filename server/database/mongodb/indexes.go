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
	ensureUserIndexes(db, dbName)
	ensureVaultIndexes(db, dbName)
	ensurePasswordIndexes(db, dbName)
}

// -------------------------
// User Collection Indexes
// -------------------------
func ensureUserIndexes(db *mongo.Client, dbName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userCol := db.Database(dbName).Collection(UserCollection)

	// Unique index on username
	_, err := userCol.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		log.Println("Could not create user index:", err)
	}
}

// -------------------------
// Vault Collection Indexes
// -------------------------
func ensureVaultIndexes(db *mongo.Client, dbName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	vaultCol := db.Database(dbName).Collection(VaultCollection)

	// Index on user_id for fetching all vaults of a user
	_, err := vaultCol.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "user_id", Value: 1}},
	})
	if err != nil {
		log.Println("Could not create vault user_id index:", err)
	}

	// Unique compound index on user_id + title
	_, err = vaultCol.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "user_id", Value: 1}, {Key: "title", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		log.Println("Could not create vault compound index (user_id + title):", err)
	}

	// Index to sort vaults by creation date per user
	_, err = vaultCol.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "user_id", Value: 1}, {Key: "created_at", Value: -1}},
	})
	if err != nil {
		log.Println("Could not create vault sort index:", err)
	}
}

// -------------------------
// Password Collection Indexes
// -------------------------
func ensurePasswordIndexes(db *mongo.Client, dbName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	passwordCol := db.Database(dbName).Collection(PasswordCollection)

	// Index on vault_id
	_, err := passwordCol.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "vault_id", Value: 1}},
	})
	if err != nil {
		log.Println("Could not create password vault_id index:", err)
	}

	// Unique compound index on vault_id + title
	_, err = passwordCol.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "vault_id", Value: 1}, {Key: "title", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		log.Println("Could not create password compound index (vault_id + title):", err)
	}

	// Index on vault_id + website for search
	_, err = passwordCol.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "vault_id", Value: 1}, {Key: "website", Value: 1}},
	})
	if err != nil {
		log.Println("Could not create password vault_id + website index:", err)
	}

	// Index to sort passwords by creation date per vault
	_, err = passwordCol.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "vault_id", Value: 1}, {Key: "created_at", Value: -1}},
	})
	if err != nil {
		log.Println("Could not create password vault_id + created_at sort index:", err)
	}
}
