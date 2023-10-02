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


func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "*")

        // Handle preflight requests
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}
func main() {
	r := mux.NewRouter()
	r.Use(corsMiddleware)
	r.HandleFunc("/image/{name}", api.HandleFunc(handlers.FileHandler))
	r.HandleFunc("/health", api.HandleFunc(handlers.ServerHealthHandler))
	r.HandleFunc("/arcs", api.HandleFunc(handlers.GetAllArcsHandler))
	r.HandleFunc("/auth/github/{code}", api.HandleFunc(handlers.AuthWithGithubHandler))

	color := color.New(color.FgHiCyan).Add(color.Underline)
	color.Println("🏮 Server Started and Listening on port", serverPort)
	log.Fatal(http.ListenAndServe(serverPort, r))
}
