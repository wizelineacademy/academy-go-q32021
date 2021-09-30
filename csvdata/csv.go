package csvdata

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/Diegoplas/go-bootcamp-deliverable/config"
	"github.com/Diegoplas/go-bootcamp-deliverable/model"
)

func ListPokemons() ([]model.PokemonData, error) {

	// open the file
	csvFile, err := os.Open(config.CSVPath)
	if err != nil {
		return nil, fmt.Errorf("error opening csv file: %v", err.Error())
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading csv file: %v", err.Error())
	}

	allPokemons, err := linesToSlice(csvLines)
	if err != nil {
		return nil, fmt.Errorf("data handling error: %v", err.Error())
	}

	return allPokemons, nil
}

func linesToSlice(csvLines [][]string) ([]model.PokemonData, error) {

	var allPokemons []model.PokemonData

	for _, line := range csvLines {

		pokemonID, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, fmt.Errorf("error converting string to int %v", err.Error())
		}

		if line[3] == "" {
			line[3] = " - "
		}

		legendary := false
		if boolVal, err := strconv.ParseBool(line[4]); err == nil {
			legendary = boolVal
		} else {
			return nil, fmt.Errorf("error parsing bool %v", err.Error())
		}

		pokemon := model.PokemonData{
			ID:        pokemonID,
			Name:      line[1],
			Type1:     line[2],
			Type2:     line[3],
			Legendary: legendary,
		}
		allPokemons = append(allPokemons, pokemon)
	}

	return allPokemons, nil
}
