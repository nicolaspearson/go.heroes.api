package controllers

import (
	"encoding/json"
	"net/http"

	"go-hero/models"
	u "go-hero/utils"
)

// CreateHero : Creates a new hero
var CreateHero = func(w http.ResponseWriter, r *http.Request) {
	hero := &models.Hero{}
	// Decode the request body into a struct
	err := json.NewDecoder(r.Body).Decode(hero)
	if err != nil {
		// Decoding failed, return an error
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	// Create the hero
	resp := hero.Create()
	u.Respond(w, resp)
}
