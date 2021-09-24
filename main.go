package main

import (
	"log"
	"net/http"

	"github.com/Marcxz/academy-go-q32021/controller"
	"github.com/gorilla/mux"
)

var (
	ac = controller.NewAddressController()
)

func main() {

	r := mux.NewRouter()
	p := ":3000"
 
	
	r.HandleFunc("/address", ac.ReadCSVAddress).Methods("GET")

	log.Printf("API listening on Port%s", p)
	
	log.Fatalln(http.ListenAndServe(p, r))
}
