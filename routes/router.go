package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"wizeline/controller"
)

func SetupRouter(router *mux.Router) *mux.Router{
	router = mux.NewRouter()
	router.HandleFunc("/", controller.HomeHandler).Methods(http.MethodGet)
	router.HandleFunc("/status", controller.StatusHandler).Methods(http.MethodGet)
	router.HandleFunc("/users", controller.UserHandler).Methods(http.MethodGet)
	//router.HandleFunc("/users/{userId}", controller.UserFilterByIdHandler).Methods(http.MethodGet)
	return  router
}

