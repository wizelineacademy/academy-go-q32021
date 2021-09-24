package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router = router.PathPrefix("/api/v1/").Subrouter()

	router.HandleFunc("/definitions/", func(rw http.ResponseWriter, r *http.Request) {}).
		Queries("term", "{term}").
		Methods(http.MethodGet)

	router.HandleFunc("/definitions/csv/", func(rw http.ResponseWriter, r *http.Request) {}).
		Queries("concurrent", "{concurrent:^(true|false)}").
		Methods(http.MethodGet)
	return mux.NewRouter()
}
