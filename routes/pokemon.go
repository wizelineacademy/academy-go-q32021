package routes

import (
	"github.com/MoraAlex/academy-go-q32021/controller"

	"github.com/gorilla/mux"
)

//Get handler routes
func Get(router *mux.Router) {
	router.HandleFunc("/pokemons", controller.GetAllPokemons)
	router.HandleFunc("/pokemons/{id}", controller.GetPokemonById)
}
