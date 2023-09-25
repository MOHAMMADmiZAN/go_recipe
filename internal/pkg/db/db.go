package db

import (
	"context"
	"errors"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var client *mongo.Client

// Init initializes the MongoDB client and sets up mgm configuration.
func Init() error {
	if client != nil {
		return nil // Already initialized
	}

	connectionURL := os.Getenv("DB_CONNECTION_URL")
	dbName := os.Getenv("DB_NAME")
	userName := os.Getenv("MONGO_ROOT_USERNAME")
	password := os.Getenv("MONGO_ROOT_PASSWORD")
	credentials := options.Credential{
		Username: userName,
		Password: password,
	}

	clientOptions := options.Client().ApplyURI(connectionURL).SetAuth(credentials)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	if err := mgm.SetDefaultConfig(nil, dbName, clientOptions); err != nil {
		return err
	}

	log.Println("Database connected")
	return nil
}

// GetClient returns a reference to the MongoDB client.
func GetClient() (*mongo.Client, error) {
	if client == nil {
		return nil, errors.New("database client is nil")
	}
	return client, nil
}

// Disconnect disconnects the MongoDB client.
func Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if client == nil {
		return errors.New("database client is nil")
	}
	log.Print("Disconnecting from database...")
	return client.Disconnect(ctx)
}
