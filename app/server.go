package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"wizeline/routes"
)

var Router *mux.Router

func Start(){
	Router = routes.SetupRouter(Router)
	if Router != nil {
		http.Handle("/", Router)
		log.Println("Server is running")
		log.Fatal(http.ListenAndServe(":8080", Router))
	}else{
		log.Println("Error setting server routes")
	}
}