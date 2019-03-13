package main

import (
	"fmt"
	"net/http"
	"os"

	"go-hero/app"
	"go-hero/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/hero", controllers.CreateHero).Methods("POST")
	router.HandleFunc("/hero/{id}", controllers.UpdateHero).Methods("PUT")
	router.HandleFunc("/hero", controllers.GetHero).Methods("GET")
	router.HandleFunc("/heroes", controllers.GetHeroes).Methods("GET")

	// Attach the JWT auth middleware
	router.Use(app.JwtAuthentication)

	// Get port from .env file
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	// Launch the app, visit http://localhost:8000/api
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
