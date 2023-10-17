package handlers

import (
	"fmt"
	"net/http"

	"github.com/EvandrooViegas/api"
)

func ServerHealthHandler(w http.ResponseWriter, r *http.Request) error  {
	fmt.Println(r.Host)
	return api.HandleJSONResponse(w, api.ApiResponse{
		Status: http.StatusOK,
		Message: "Server is alive",
	})
}