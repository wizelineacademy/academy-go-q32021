package usecase

import (
	"errors"
	"fmt"

	"github.com/Diegoplas/go-bootcamp-deliverable/csvdata"
	"github.com/Diegoplas/go-bootcamp-deliverable/model"
)

func GetPokemonFromCSV(wantedIndex int) (model.PokemonData, error) {

	allPokemons, err := csvdata.ListPokemons()
	if err != nil {
		return model.PokemonData{}, fmt.Errorf("error reading CSV file: %v", err.Error())
	}

	for _, pokemon := range allPokemons {
		if pokemon.ID == wantedIndex {
			return pokemon, nil
		}
	}
	return model.PokemonData{}, errors.New("no pokemon found")
}
