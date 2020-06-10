package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetDatabase :Return database
func GetDatabase() (*mongo.Database, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://huyhnc81:Huysamdua1@jusay-gd13q.gcp.mongodb.net/test?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	return client.Database("jusay"), nil
}
