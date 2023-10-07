package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/EvandrooViegas/api"
	"github.com/gofor-little/env"
	"github.com/gorilla/mux"
)

func AuthWithGithubHandler(w http.ResponseWriter, r *http.Request) error {
	pathnames := mux.Vars(r)
	code, ok := pathnames["code"]
	if !ok {
		return fmt.Errorf("a code pathname was not provided")
	}
	gitID := loadEnvVariable("GITHUB_CLIENT_ID")
	gitScrt := loadEnvVariable("GITHUB_CLIENT_SECRET")

	
	acsTkn, err := getGHaccesssTkn(map[string]string{
		"client_id": gitID,
		"client_secret": gitScrt,
		"code": code,
	})
	if err != nil {
		return fmt.Errorf("Error parsing")
	}

	
	usrRespMap, err := getGHuser(acsTkn)
	if err != nil {
		return err
	}
	fmt.Println(usrRespMap["id"])
	return nil
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

func getGHaccesssTkn(data map[string]string) (string, error) {
	acsTknMap, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	acsTknReq, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(acsTknMap),
	)
	if err != nil {
		return "", err
	}
	acsTknReq.Header.Set("Content-Type", "application/json")
	acsTknReq.Header.Set("Accept", "application/json")
	var acsTknResMap map[string]interface{}
	api.MakeRequest(acsTknReq, &acsTknResMap)
	result, ok := acsTknResMap["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("Error parsing string")
	}
	return result, nil
}

func loadEnvVariable(key string) string {
	appEnv := os.Getenv("APP_ENV")
	switch appEnv {
		case "dev":
		if err := env.Load(".env.local"); err != nil {
			return ""
		}
		default:
		if err := env.Load(".env.prod"); err != nil {
			return ""
		}
	}
	return env.Get(key, "")
}
