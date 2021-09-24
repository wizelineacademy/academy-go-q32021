package infraestructure

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Router - Interface to mock the router interface
type Router interface {
	Get(uri string, f func(http.ResponseWriter, *http.Request))
	Post(uri string, f func(http.ResponseWriter, *http.Request))
	Serve(p string)
}

var (
	md = mux.NewRouter()
)

type muxRouter struct{}

//NewMuxRouter - like the constructor of the Router to handle all the request from the user
func NewMuxRouter() Router {
	return &muxRouter{}
}

// Get - Refactor and handle the get request from the user
func (*muxRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	md.HandleFunc(uri, f).Methods("GET")
}

// Post - Refactor and handle the post request from the user
func (*muxRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	md.HandleFunc(uri, f).Methods("POST")
}

// Server - Up and run the project
func (*muxRouter) Serve(p string) {
	fmt.Printf("Server is running on port %s", p)
	http.ListenAndServe(p, md)
}
