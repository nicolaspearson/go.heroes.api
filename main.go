package main

import (
	"fmt"
	"net/http"
	"os"

	"go-hero/app"
	"go-hero/src/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	// Attach JWT auth middleware
	router.Use(app.JwtAuthentication)

	// Get port from .env file, we did not specify any port
	// so this should return an empty string when tested locally
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	// Launch the app, visit localhost:8000/api
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
