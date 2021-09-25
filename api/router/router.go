package router

import (
	c "api/controller"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var Resp = render.New()

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api", c.HelloWorld)
	router.HandleFunc("/api/import", c.ReadCSV)

	return router
}
