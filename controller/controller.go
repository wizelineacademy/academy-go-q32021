package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Diegoplas/go-bootcamp-deliverable/model"
	"github.com/Diegoplas/go-bootcamp-deliverable/usecase"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func GetPokemon(w http.ResponseWriter, r *http.Request) {

	requestIndex := mux.Vars(r)["id"]

	wantedIndex, err := validateID(requestIndex)
	if err != nil {
		errorResponse := model.ErrorResponse{Err: err.Error()}
		render.New().JSON(w, http.StatusBadRequest, errorResponse)
		return
	}

	response, err := usecase.GetPokemonFromCSV(wantedIndex)
	if err != nil {
		errorResponse := model.ErrorResponse{Err: err.Error()}
		render.New().JSON(w, http.StatusInternalServerError, errorResponse)
		return
	}

	render.New().JSON(w, http.StatusOK, &response)
}

func validateID(index string) (int, error) {

	wantedIndex, err := strconv.Atoi(index)
	if err != nil {
		return 0, fmt.Errorf("string to int convertion failed %v", err.Error())
	}

	if wantedIndex < 1 || wantedIndex > 151 {
		log.Println("Please introduce a valid pokemon ID from first gen. (1-151)")
		return 0, errors.New("invalid id")
	}

	return wantedIndex, nil
}
