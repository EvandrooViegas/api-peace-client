package services

import (
	"context"
	"fmt"
	"time"

	"github.com/EvandrooViegas/db"
	"github.com/EvandrooViegas/types"
	"github.com/EvandrooViegas/utils"

)

const TOKEN_AND_COOKIE_EXPIRATION = (time.Hour * 24) * 30 // 1 month

func AuthWithGithub(code string) (string, error) {
	gitID, err := utils.LoadEnvVariable("GITHUB_CLIENT_ID")
	if err != nil {
		return "", err
	}
	gitScrt, err := utils.LoadEnvVariable("GITHUB_CLIENT_SECRET")
	if err != nil {
		return "", err
	}
	acsTkn, err := getGHaccesssTkn(map[string]string{
		"client_id":     gitID,
		"client_secret": gitScrt,
		"code":          code,
	})
	if err != nil {
		return "", err
	}
	usrRespMap, err := getGHuser(acsTkn)
	if err != nil {
		return "", err
	}
	id, ok := usrRespMap["id"]
	if !ok {
		return "", fmt.Errorf("Could not get user id from the response map")
	}
	id, ok = id.(float64)
	if !ok {
		return "", fmt.Errorf("Could not transform id into a float64")
	}

	
	mongo, disc, err := db.ConnectMongoDB()
	if err != nil {
		return "", err
	}
	defer disc(context.TODO())
	usrService := mongo.GetUserService()


	exists, foundUsr, err := usrService.GetByGitHubID(id.(float64))
	if err != nil {
		return "", err
	}

	var usrID string
	exists = false
	if !exists {
		providerID, ok := usrRespMap["id"].(float64)
		if !ok {
			return "", fmt.Errorf("Error parsing the response user id to float64")
		}
		nUser, err := usrService.InsertUser(types.NewUser{
			AvatarURL:  usrRespMap["avatar_url"].(string),
			Username:   usrRespMap["login"].(string),
			ProviderID: providerID,
			Provider:   "github",
		})
		if err != nil {
			return "", err
		}
		usrID = nUser.ID
		fmt.Println("nUser: ", nUser)
	} else {
		usrID = foundUsr.ID
		fmt.Println("foundUsr: ", foundUsr)

	}
	fmt.Println(usrID)
	tkn, err := createAuthJWT(usrID)
	if err != nil {
		return "", err
	}
	return tkn, nil
}


