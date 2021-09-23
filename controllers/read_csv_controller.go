package controllers

import (
	services "academy-go-q32021/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadCsv(w http.ResponseWriter, r *http.Request, id string) {

	p, err := services.ReadCsv(id)

	fmt.Println("controller reading err", err)
	if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	pokemonJson, _ := json.Marshal(p)
	w.Write(pokemonJson)
}
