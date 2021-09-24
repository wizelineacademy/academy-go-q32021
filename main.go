package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Marcxz/academy-go-q32021/controller"
)

var (
	ac = controller.NewAddressController()
)

func main() {

	r := mux.NewRouter()
	p := ":3000"

	r.HandleFunc("/address", ac.NewAddressController).Methods("GET")

	log.Printf("API listening on Port%s", p)
	log.Fatalln(http.ListenAndServe(p, r))
}
