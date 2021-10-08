package route

import (
	"net/http"

	"github.com/Diegoplas/go-bootcamp-deliverable/controller"

	"github.com/gorilla/mux"
)

func GetRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/pokemon/{id}", controller.GetPokemonHandler).Methods(http.MethodGet)
	return
}
