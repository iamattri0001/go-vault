package mongodb

import (
	"context"
	"fmt"
	"go-vault/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConnection struct {
	client       *mongo.Client
	dbName       string
	QueryTimeout time.Duration
}

func NewMongoDbConnection(config *config.MongoConfig) (*MongoDBConnection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.ConnectionTimeout)*time.Millisecond)
	defer cancel()

	clientOpts := options.Client().
		ApplyURI(config.URI).
		SetMaxPoolSize(config.MaxPoolSize).
		SetMinPoolSize(config.MinPoolSize).
		SetMaxConnIdleTime(time.Duration(config.IdleConnTimeout) * time.Millisecond).
		SetServerSelectionTimeout(time.Duration(config.ServerSelectionTimeout) * time.Millisecond).
		SetConnectTimeout(time.Duration(config.ConnectTimeout) * time.Millisecond).
		SetSocketTimeout(time.Duration(config.SocketTimeout) * time.Millisecond).
		SetRetryWrites(config.RetryWrites).
		SetRetryReads(config.RetryReads)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	// Ping the database to check if the connection string is valid and the server is reachable
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	ensureIndexes(client, config.Database)

	dbConnection := &MongoDBConnection{
		client:       client,
		dbName:       config.Database,
		QueryTimeout: time.Duration(config.QueryTimeout) * time.Millisecond,
	}

	return dbConnection, nil
}

func (c *MongoDBConnection) GetDatabase() *mongo.Database {
	return c.client.Database(c.dbName)
}
