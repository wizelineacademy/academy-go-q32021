package controller

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"wizeline/usecase"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users , err := usecase.GetAllUsers()
	if err == nil {
		w.WriteHeader(http.StatusOK)
		_ ,err := fmt.Fprintf(w, users)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			http.Error(w,"Error getting users",http.StatusInternalServerError)
		}
	}else{
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		http.Error(w,"Error reading users csv file",http.StatusInternalServerError)
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Home")
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"Running": true}`)
}