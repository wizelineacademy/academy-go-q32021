package app

import (
	"github.com/bifr0ns/academy-go-q32021/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	routes.Setup(router)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
