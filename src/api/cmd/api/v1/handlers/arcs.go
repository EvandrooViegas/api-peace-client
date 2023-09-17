package handlers

import (
	"net/http"

	"github.com/EvandrooViegas/api-piece/cmd/api/v1/api"
	"github.com/EvandrooViegas/api-piece/cmd/api/v1/services"
)


func GetAllArcsHandler(w http.ResponseWriter, r *http.Request) error {
	addr := api.GetServerAddr(r)
	arcs, err := services.GetAllArcs(addr)
	if err != nil {
		return api.HandleJSONResponse(w, api.ApiResponse{
			Status: http.StatusInternalServerError,
			Message: "Fetched Arcs Successfully",
			Error: err,
		})
	}
	return api.HandleJSONResponse(w, api.ApiResponse{
		Status: http.StatusOK,
		Data: arcs,
		Message: "Fetched Arcs Successfully",
	})
}