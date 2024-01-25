package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// NewMongoDB creates a new MongoDB client and establishes a connection.
func NewMongoDB(uri, dbName string) (*MongoDB, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to create new MongoDB client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the primary to ensure network connectivity
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	db := client.Database(dbName)

	return &MongoDB{Client: client, Database: db}, nil
}

// GetCollection returns a handle to a collection in the database.
func (m *MongoDB) GetCollection(collectionName string) *mongo.Collection {
	return m.Database.Collection(collectionName)
}

// CloseConnection closes the MongoDB connection.
func (m *MongoDB) CloseConnection() {
	if m.Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := m.Client.Disconnect(ctx); err != nil {
			log.Printf("Warning: failed to close MongoDB connection: %v", err)
		}
	}
}
