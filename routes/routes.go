package routes

import (
	"github.com/gorilla/mux"
	"github.com/s1nuh3/academy-go-q32021/controllers"
)

// Get - Add the handlers to the get methods aviable
func Get(router *mux.Router) {
	router.HandleFunc("/", controllers.IndexHandler()).Methods("GET")
	router.HandleFunc("/users", controllers.GetUsers()).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUsersbyId()).Methods("GET")
}
