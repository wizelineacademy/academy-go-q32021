package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MoraAlex/academy-go-q32021/services"

	"github.com/gorilla/mux"
)

// GetALlPokemons json repond to get All pokemons
func GetAllPokemons(w http.ResponseWriter, r *http.Request) {
	pokemons, err := services.GetAllPokemons()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(pokemons)
}

// GetPokemonById json repond to get a pokemon by ID
func GetPokemonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	pokemons, err := services.GetPokemonById(id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(pokemons)

}
