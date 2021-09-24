package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bifr0ns/academy-go-q32021/services"
	"github.com/gorilla/mux"
)

func GetPokemonById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pokemonId := vars["pokemon_id"]

	pokemon, httpStatus, err := services.GetPokemonInfo(pokemonId)

	rw.Header().Add("Content-Type", "application/json")

	if err != nil {
		if httpStatus == 404 {
			rw.WriteHeader(http.StatusNotFound)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(rw).Encode(err.Error())
		return
	}

	json.NewEncoder(rw).Encode(pokemon)
}
