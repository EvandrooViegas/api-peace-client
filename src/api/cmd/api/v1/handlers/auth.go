package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/EvandrooViegas/api"
	"github.com/EvandrooViegas/services"
	"github.com/gorilla/mux"
)

func AuthWithGithubHandler(w http.ResponseWriter, r *http.Request) error {
	pathnames := mux.Vars(r)
	code, ok := pathnames["code"]

	if !ok {
		return fmt.Errorf("a code pathname was not provided")
	}
	tkn, err := services.AuthWithGithub(code)
	if err != nil {
		return api.HandleJSONResponse(w, api.ApiResponse{
			Status: http.StatusInternalServerError,
			Error:  err,
		})
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "auth-jwt",
		Value:   tkn,
		Expires: time.Now().Add(services.TOKEN_AND_COOKIE_EXPIRATION),
		Path: "/",
	})

	return api.HandleJSONResponse(w, api.ApiResponse{
		Status: http.StatusOK,
		Data: "created",
		Error:  err,
	})
}
