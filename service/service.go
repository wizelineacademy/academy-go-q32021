package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Diegoplas/go-bootcamp-deliverable/model"
)

func GetPokemonFromExternalAPI(wantedIndex string) (model.PokemonData, error) {

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s?name", wantedIndex)

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	data, err := ioutil.ReadAll(response.Body) // manage error
	if err != nil {
		return model.PokemonData{}, fmt.Errorf("error reading response body %v", err.Error())
	}

	pokemonData := model.PokemonExternalData{}
	unmarshalErr := json.Unmarshal(data, &pokemonData)
	if unmarshalErr != nil {
		return model.PokemonData{}, fmt.Errorf("error performing the unmarshal %v", err.Error())
	}
	//loggear el error
	//controller se encarga de distintas situaciones.

	formatedPokemonData := formatPokemonData(pokemonData)

	return formatedPokemonData, nil
}

func formatPokemonData(externalPokemonData model.PokemonExternalData) model.PokemonData {

	formatedPokemonData := model.PokemonData{}

	if len(externalPokemonData.Types) == 1 {
		formatedPokemonData = model.PokemonData{
			ID:     externalPokemonData.ID,
			Name:   externalPokemonData.Name,
			Height: externalPokemonData.Height,
			Type1:  externalPokemonData.Types[0].Type.Name,
			Type2:  " - ",
		}
	} else if len(externalPokemonData.Types) == 2 {
		formatedPokemonData = model.PokemonData{
			ID:     externalPokemonData.ID,
			Name:   externalPokemonData.Name,
			Height: externalPokemonData.Height,
			Type1:  externalPokemonData.Types[0].Type.Name,
			Type2:  externalPokemonData.Types[1].Type.Name,
		}
	}

	return formatedPokemonData
}
