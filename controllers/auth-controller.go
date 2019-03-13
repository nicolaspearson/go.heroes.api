package controllers

import (
	"encoding/json"
	"net/http"

	"go-hero/models"
	u "go-hero/utils"
)

// Authenticate : Authenticates a user
var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	// Decode the request body into a struct
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// Decoding failed, return an error
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(user.Email, user.Password)
	u.Respond(w, resp)
}

// RegisterUser : Registers a new user
var RegisterUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	// Decode the request body into a struct
	if err != nil {
		// Decoding failed, return an error
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	// Create the user
	resp := user.Create()
	u.Respond(w, resp)
}
