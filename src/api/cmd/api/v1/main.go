package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EvandrooViegas/api-piece/cmd/api/v1/api"
	"github.com/EvandrooViegas/api-piece/cmd/api/v1/handlers"
	"github.com/gorilla/mux"
)

const PORT = 8080
var serverPort string = fmt.Sprint(":", PORT)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", api.HandleFunc(handlers.ServerHealthHandler))
	r.HandleFunc("/arcs", api.HandleFunc(handlers.GetAllArcsHandler))
	log.Fatal(http.ListenAndServe(serverPort, r))
}
