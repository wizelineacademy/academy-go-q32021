package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/api/v1/hello", hello)
	myRouter.HandleFunc("/api/v1/csv", csv)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func hello(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"response": "hello",
	}
	json.NewEncoder(w).Encode(response)
}

func csv(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
