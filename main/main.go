package main

import (
	"log"
	"net/http"

	"github.com/Diegoplas/go-bootcamp-deliverable/config"
	"github.com/Diegoplas/go-bootcamp-deliverable/route"

	"github.com/gorilla/handlers"
)

func main() {
	router := route.GetRouter()
	methods := handlers.AllowedMethods([]string{http.MethodGet})
	log.Fatal(http.ListenAndServe(config.Port, handlers.CORS(methods)(router)))
}
