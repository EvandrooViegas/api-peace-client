package handlers

import (
	"net/http"

	"github.com/EvandrooViegas/api-piece/cmd/api/v1/api"
)

func ServerHealthHandler(w http.ResponseWriter, r *http.Request) error  {
	return api.HandleJSONResponse(w, api.ApiResponse{
		Status: http.StatusOK,
		Message: "Server is alive",
	})
}