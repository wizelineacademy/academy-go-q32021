package routes

import (
	"github.com/Diegoplas/go-bootcamp-deliverable/controllers"
	"github.com/gorilla/mux"
)

func GetRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/pokedex/{pokemon_id}", controllers.GetPokemon).Methods("GET")
	return
}
