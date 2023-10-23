package services

import (
	"context"
	"fmt"

	"github.com/EvandrooViegas/db"
	"github.com/EvandrooViegas/types"
)

func GetUserByID(id string) (bool, types.User, error) {
	m, disc, err := db.ConnectMongoDB()
	defer disc(context.TODO())
	if err != nil {
		return false, types.User{}, err
	}
	usrService := m.GetUserService()

	exists, usr, err := usrService.GetByID(id)
	if err != nil {
		return false, types.User{}, err
	}
	if !exists {
		return false, types.User{}, fmt.Errorf("user not found")
	}
	return true, usr, nil
}