package router

import (
	c "github.com/emamex98/academy-go-q32021/controller"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var Resp = render.New()

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router = router.PathPrefix("/api").Subrouter()

	router.HandleFunc("/", c.HelloWorld)
	router.HandleFunc("/contestants", c.GetContestans)
	router.HandleFunc("/contestants/{id}", c.GetSingleContestant)

	return router
}
