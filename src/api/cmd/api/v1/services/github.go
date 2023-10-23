package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/EvandrooViegas/api"
)

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

	accTkn, ok := res["access_token"]
	if !ok {
		return "", fmt.Errorf("Could not get  access_token")
	}
	return accTkn.(string), nil

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
