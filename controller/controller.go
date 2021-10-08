package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Diegoplas/go-bootcamp-deliverable/csvdata"
	"github.com/Diegoplas/go-bootcamp-deliverable/model"
	"github.com/Diegoplas/go-bootcamp-deliverable/service"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func GetPokemonHandler(w http.ResponseWriter, r *http.Request) {

	requestIndex := mux.Vars(r)["id"]
	validatedIndex, err := validateID(requestIndex)
	if err != nil {
		errorResponse := model.ErrorResponse{Err: err.Error()}
		render.New().JSON(w, http.StatusBadRequest, errorResponse)
		return
	}
	index := strconv.Itoa(validatedIndex)

	existingPokemon, err := csvdata.CheckIfPokemonAlreadyExists(index)
	if err != nil {
		errorResponse := model.ErrorResponse{Err: err.Error()}
		render.New().JSON(w, http.StatusInternalServerError, errorResponse)
		return
	}

	var response model.PokemonData
	// If pokemon is an empty model, get it from external api.
	if existingPokemon == (model.PokemonData{}) {
		response, err = service.GetPokemonFromExternalAPI(index)
		if err != nil {
			log.Printf("Error on GetPokemonFromExternalAPI %s", err.Error())
			errorResponse := model.ErrorResponse{Err: err.Error()}
			render.New().JSON(w, http.StatusInternalServerError, errorResponse)
			return
		}
		err = csvdata.WritePokemonIntoCSV(response)
		if err != nil {
			log.Printf("Error on WritePokemonIntoCSV: %s", err.Error())
			errorResponse := model.ErrorResponse{Err: err.Error()}
			render.New().JSON(w, http.StatusInternalServerError, errorResponse)
			return
		}

	} else {
		// Get it from CSV file instead.
		response = model.PokemonData{
			ID:     existingPokemon.ID,
			Name:   existingPokemon.Name,
			Height: existingPokemon.Height,
			Type1:  existingPokemon.Type1,
			Type2:  existingPokemon.Type2,
		}
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
