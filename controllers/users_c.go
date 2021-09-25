package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/s1nuh3/academy-go-q32021/services"
)

// GetUsers - Returns the list of users
func GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := services.GetUsersfromCSV()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error :"+err.Error())
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(u)
	}
}

// GetUsersbyId - Look up for a user id
func GetUsersbyId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		u, err := services.GetUserbyIdfromCSV(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error :"+err.Error())
		}
		if u.Id == 0 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(u)
		}

	}
}
