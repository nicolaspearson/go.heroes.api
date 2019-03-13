package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

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

// GetHero : Fetches a the hero associated with the provided id
var GetHero = func(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idParam, ok := params["id"]
	if !ok {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	id, err := strconv.Atoi(idParam[0])
	if err != nil {
		// The parameter is not an integer
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	data := models.GetOne(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetHeroes : Fetches all the heroes
var GetHeroes = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetAll()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
