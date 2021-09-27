package controllers

import (
	"net/http"

	"github.com/Diegoplas/go-bootcamp-deliverable/app"
	"github.com/Diegoplas/go-bootcamp-deliverable/useCase"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func GetPokemon(w http.ResponseWriter, r *http.Request) {

	requestIndex := mux.Vars(r)["pokemon_id"]

	wantedIndex := app.ValidateID(requestIndex)

	response := useCase.GetPokemonFromCSV(wantedIndex)

	render.New().JSON(w, http.StatusOK, &response)
}
