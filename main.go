package main

import (
	"log"
	"net/http"

	"github.com/MoraAlex/academy-go-q32021/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	routes.Get(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
