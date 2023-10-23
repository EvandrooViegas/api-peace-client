package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/EvandrooViegas/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userService struct {
	client  *mongo.Client
}


func (u userService) GetByID(id string) (bool, types.User, error) {
	coll :=  u.client.Database(MONGO_DATABASE).Collection(USER_COLLECTION)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, types.User{}, err
	}
	q := bson.D{{ "_id", objID}}

	var res types.User
	if err := coll.FindOne(context.TODO(), q).Decode(&res); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, types.User{}, nil
		}
		return false, types.User{}, err
	}
	res.ID = id
	return true, res, nil
}
func (u userService) GetByGitHubID(id float64) (bool, types.User, error) {
	coll :=  u.client.Database(MONGO_DATABASE).Collection(USER_COLLECTION)
	q := bson.D{{ "provider_id", id}, {"provider", "github"}}

	var res types.User
	if err := coll.FindOne(context.TODO(), q).Decode(&res); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, types.User{}, nil
		}
		return false, types.User{}, err
	}
	return true, res, nil
}


func (u userService) InsertUser(nUser types.NewUser) (types.User, error) {
	coll := u.client.Database(MONGO_DATABASE).Collection(USER_COLLECTION)
	res, err := coll.InsertOne(context.TODO(), nUser)
	if err != nil {
		return types.User{}, err 
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return types.User{
		ID: id,
		AvatarURL: nUser.AvatarURL,
		Username: nUser.Username,
		ProviderID: nUser.ProviderID,
		Provider: nUser.Provider,
	}, nil
}