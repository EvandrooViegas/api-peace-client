package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiResponse struct {
	Status  int
	Message string
	Data    interface{}
	Error   interface{}
}

func HandleJSONResponse(w http.ResponseWriter, r ApiResponse) error {
	data, err := json.Marshal(map[string]interface{}{
		"status": r.Status,
		"message": r.Message,
		"data": r.Data,
	}) 
	if err != nil {
		return err
	}
	if r.Error != nil {
		fmt.Println("Error: ", r.Error)
	}
	
	w.WriteHeader(r.Status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	return nil
}

func HandleFunc(
	handler func(w http.ResponseWriter, r *http.Request) error,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			fmt.Println("Handler Error: ", err)
		}
	}
}