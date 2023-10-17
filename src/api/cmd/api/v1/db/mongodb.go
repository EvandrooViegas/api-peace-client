package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://mongodb:27017"
const MONGO_DATABASE = "oneapi" 
const USER_COLLECTION = "users" 
type MongoDB struct {
	client *mongo.Client
}

func (cl MongoDB) GetUserService() userService {
	return userService{ client: cl.client }
}

func (cl MongoDB) Disconnect(ctx context.Context) error {
	if err := cl.client.Disconnect(ctx); err != nil {
		return err
	}

	return nil
}

func ConnectMongoDB() (MongoDB, error) {
	ctx := context.TODO()
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return MongoDB{}, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return MongoDB{}, err
	}
	fmt.Println("Connected to the database")

	return MongoDB{
		client: client,
	}, nil
}
