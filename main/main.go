package main

import (
	"log"
	"net/http"

	"github.com/Diegoplas/go-bootcamp-deliverable/routes"
	"github.com/gorilla/handlers"
)

func main() {
	router := routes.GetRouter()
	methods := handlers.AllowedMethods([]string{"GET"})
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(methods)(router)))
}
