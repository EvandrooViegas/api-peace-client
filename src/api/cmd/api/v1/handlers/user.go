package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/EvandrooViegas/api"
	"github.com/EvandrooViegas/services"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) error {
	authHeader, exists := r.Header["Authorization"]
	if !exists {
		return api.HandleJSONResponse(w, api.ApiResponse{
			Status: http.StatusBadRequest,
			Error: fmt.Errorf("No auth header provided"),
		})
	}
	tkn := strings.Split(authHeader[0], "Bearer ")[1]
	id, err := services.ReadPlayerToken(tkn)
	if err != nil {
		return api.HandleJSONResponse(w, api.ApiResponse{
			Status: http.StatusInternalServerError,
			Error: err,
		})
	}
	fmt.Println("id: ", id)

	exists, usr, err := services.GetUserByID(id)
	if err != nil {
		return api.HandleJSONResponse(w, api.ApiResponse{
			Status: http.StatusInternalServerError,
			Error: err,
		})
	}

	if !exists {
		return api.HandleJSONResponse(w, api.ApiResponse{
			Status: http.StatusOK,
			Message: "user was not found",
			Error: fmt.Errorf("user was not found"),
		})
	}


	return api.HandleJSONResponse(w, api.ApiResponse{
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"user": usr,
		},
	})
}
