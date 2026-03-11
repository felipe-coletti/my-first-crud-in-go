package mongodb

import (
	"context"

	"github.com/felipe-coletti/my-first-crud-in-go/src/config/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InitMongoDB() {
	ctx := context.Background()

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	logger.Info("Successfully connected")
}
