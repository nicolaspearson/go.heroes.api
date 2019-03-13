package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-hero/models"
	u "go-hero/utils"

	"github.com/gorilla/mux"
)

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

	data := models.GetHero(uint(id))
	if data == nil {
		u.Respond(w, u.Message(false, "Record not found!"))
		return
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetHeroes : Fetches all the heroes
var GetHeroes = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetHeroes()
	if data == nil {
		u.Respond(w, u.Message(false, "No records found!"))
		return
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

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

// UpdateHero : Updates an existing hero
var UpdateHero = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	hero := &models.Hero{}
	// Decode the request body into a struct
	err = json.NewDecoder(r.Body).Decode(hero)
	if err != nil {
		// Decoding failed, return an error
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	// Update the hero
	resp := hero.Update(uint(id))
	u.Respond(w, resp)
}

// DeleteHero : Deletes the hero associated with the provided id
var DeleteHero = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	data := models.DeleteHero(uint(id))
	if data == nil {
		u.Respond(w, u.Message(false, "Record not found!"))
		return
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
