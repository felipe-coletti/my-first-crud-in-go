package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	MONGO_URI     = "MONGO_URI"
	MONGO_DB_NAME = "MONGO_DB_NAME"
)

func NewMongoDB(ctx context.Context) (*mongo.Database, error) {
	mongoURI := os.Getenv(MONGO_URI)
	mongoDBName := os.Getenv(MONGO_DB_NAME)

	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongoDBName), nil
}
