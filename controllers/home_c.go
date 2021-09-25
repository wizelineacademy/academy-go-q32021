package controllers

import (
	"fmt"
	"net/http"
)

//IndexHandler - Handles the calls to the root path of the server
func IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome, this Go Rest API is to fullfill the Wizeline Academy Go Bootcamp!!")
	}
}
