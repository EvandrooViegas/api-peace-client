package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EvandrooViegas/api"
	"github.com/EvandrooViegas/handlers"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

const PORT = 8080
var serverPort string = fmt.Sprint(":", PORT)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/image/{name}", api.HandleFunc(handlers.FileHandler))
	r.HandleFunc("/health", api.HandleFunc(handlers.ServerHealthHandler))
	r.HandleFunc("/arcs", api.HandleFunc(handlers.GetAllArcsHandler))

	c := color.New(color.FgHiCyan).Add(color.Underline)
	c.Println("🏮 Server Started and Listening on port", serverPort)
	log.Fatal(http.ListenAndServe(serverPort, r))
}
