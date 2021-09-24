package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MoraAlex/academy-go-q32021/services"
	"github.com/gorilla/mux"
)

func GetAllPokemons(w http.ResponseWriter, r *http.Request) {
	pokemons, err := services.GetAllPokemons()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(pokemons)
}

func GetPokemonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	pokemons, err := services.GetPokemonById(id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(pokemons)

}
