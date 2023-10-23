package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EvandrooViegas/api"
	"github.com/EvandrooViegas/handlers"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

var (
	port              = 8080
	serverPort string = fmt.Sprint(":", port)
)

func InitRouter() {
	r := mux.NewRouter()
	r.Use(corsMiddleware)
	r.HandleFunc("/image/{name}", api.HandleFunc(handlers.FileHandler))
	r.HandleFunc("/health", api.HandleFunc(handlers.ServerHealthHandler))
	r.HandleFunc("/arcs", api.HandleFunc(handlers.GetAllArcsHandler))
	r.HandleFunc("/auth/github/{code}", api.HandleFunc(handlers.AuthWithGithubHandler))
	r.HandleFunc("/user", api.HandleFunc(handlers.GetUserHandler))
	color := color.New(color.FgHiCyan).Add(color.Underline)
	color.Println("🏮 Server Started and Listening on port", serverPort)

	log.Fatal(http.ListenAndServe(serverPort, r))
}
