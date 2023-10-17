package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/EvandrooViegas/api"
	"github.com/EvandrooViegas/db"
	"github.com/EvandrooViegas/services"
	"github.com/EvandrooViegas/utils"
	"github.com/gorilla/mux"
)

func AuthWithGithubHandler(w http.ResponseWriter, r *http.Request) error {
	pathnames := mux.Vars(r)
	code, ok := pathnames["code"]
	if !ok {
		return fmt.Errorf("a code pathname was not provided")
	}
	gitID, err := utils.LoadEnvVariable("GITHUB_CLIENT_ID")
	if err != nil {
		return err
	}
	gitScrt, err := utils.LoadEnvVariable("GITHUB_CLIENT_SECRET")
	if err != nil {
		return err
	}

	acsTkn, err := getGHaccesssTkn(map[string]string{
		"client_id":     gitID,
		"client_secret": gitScrt,
		"code":          code,
	})
	if err != nil {
		return err
	}

	usrRespMap, err := getGHuser(acsTkn)
	if err != nil {
		return err
	}

	mongo, err := db.ConnectMongoDB()
	if err != nil {
		return err
	}
	defer mongo.Disconnect(context.TODO())

	usrService := mongo.GetUserService()
	exists, _, err := usrService.GetByGitHubID(usrRespMap["id"].(float64))
	if exists {
		fmt.Println("User already exists")
	} else {
		providerID, ok := usrRespMap["id"].(float64)
		if !ok {}
		nUser := services.NewUser{
			AvatarURL:  usrRespMap["avatar_url"].(string),
			Username:   usrRespMap["login"].(string),
			ProviderID: providerID,
			Provider:   "github",
		}

		usrService.InsertUser(nUser)

	}

	return api.HandleJSONResponse(w, api.ApiResponse{
		Status: http.StatusOK,
		Data:   usrRespMap,
	})
}

func getGHaccesssTkn(data map[string]string) (string, error) {
	requestBody, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	u, _ := url.Parse("https://github.com/login/oauth/access_token")
	q := u.Query()
	q.Set("client_id", data["client_id"])
	q.Set("client_secret", data["client_secret"])
	q.Set("code", data["code"])
	u.RawQuery = q.Encode()
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(requestBody))
	req.Header.Set("Accept", "application/json")
	if err != nil {
		return "", err
	}

	var res map[string]interface{}
	api.MakeRequest(req, &res)

	if err != nil {
		return "", err
	}

	accTkn := res["access_token"].(string)
	return accTkn, nil
}

func getGHuser(acsTkn string) (map[string]interface{}, error) {
	usrReq, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	usrReq.Header.Set("Accept", "application/vnd.github+json")
	usrReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", acsTkn))
	usrReq.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	if err != nil {
		return nil, err
	}
	var usrRespMap map[string]interface{}
	api.MakeRequest(usrReq, &usrRespMap)
	return usrRespMap, nil
}
