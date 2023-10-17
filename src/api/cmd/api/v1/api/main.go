package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ApiResponse struct {
	Status  int
	Message string
	Data    interface{}
	Error   interface{}
}


func MakeRequest(rq *http.Request, resMap *map[string]interface{}) error {
	client := &http.Client{}
	res, err := client.Do(rq)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&resMap); err != nil {
		return err
	}
	return nil
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

func HandleImageResponse(w http.ResponseWriter, buf []byte, imgName string) error {
	ext := strings.Split(imgName, ".")[1]
	contentType := "image/" + ext
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", contentType)
	w.Write(buf)
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

func GetServerAddr(r *http.Request) string {
	host := r.Host

		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}

		serverAddress := scheme + "://" + host
		return serverAddress
}