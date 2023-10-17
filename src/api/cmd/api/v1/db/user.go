package db

import (
	"context"
	"fmt"

	"github.com/EvandrooViegas/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userService struct {
	client  *mongo.Client
}

func (u userService) GetByGitHubID(id float64) (bool, services.User, error) {
	coll :=  u.client.Database(MONGO_DATABASE).Collection(USER_COLLECTION)
	q := bson.D{{ "provider_id", id}, {"provider", "github"}}

	var res services.User
	if err := coll.FindOne(context.TODO(), q).Decode(&res); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, services.User{}, nil
		}
		return false, services.User{}, err
	}
	return true, res, nil
}

func (u userService) InsertUser(nUser services.NewUser) (services.User, error) {
	coll := u.client.Database(MONGO_DATABASE).Collection(USER_COLLECTION)
	res, err := coll.InsertOne(context.TODO(), nUser)
	if err != nil {
		return services.User{}, err 
	}
	fmt.Println(res)
	return services.User{}, nil
}