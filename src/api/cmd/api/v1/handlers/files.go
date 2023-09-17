package handlers

import (
	"net/http"

	"github.com/EvandrooViegas/api-piece/cmd/api/v1/api"
	"github.com/EvandrooViegas/api-piece/cmd/api/v1/services"
	"github.com/gorilla/mux"
)

func FileHandler(w http.ResponseWriter, r *http.Request) error {
	pathnames := mux.Vars(r)
	name, ok := pathnames["name"]
	if !ok {
		return api.HandleJSONResponse(w, api.ApiResponse{
			Status: http.StatusBadRequest,
			Message: "Image name not provided",
		})
	}
	buf, err := services.GetImageBuf(name)
	if err != nil {
		return api.HandleJSONResponse(w, api.ApiResponse{
			Status: http.StatusInternalServerError,
		})
	}
	return api.HandleImageResponse(w, buf, name)
}