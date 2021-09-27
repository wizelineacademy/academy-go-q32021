package useCase

import (
	"strconv"

	"github.com/Diegoplas/go-bootcamp-deliverable/app"
	"github.com/Diegoplas/go-bootcamp-deliverable/utils"
)

func GetPokemonFromCSV(wantedIndex int) utils.PokemonData {

	csvLines := app.CSVReader()

	for _, line := range csvLines {

		pokemonID, _ := strconv.Atoi(line[0])

		if pokemonID == wantedIndex {

			if line[3] == "" {
				line[3] = " - "
			}

			legendary := false
			if line[4] == "TRUE" {
				legendary = true
			}

			pokemon := utils.PokemonData{
				ID:        pokemonID,
				Name:      line[1],
				Type1:     line[2],
				Type2:     line[3],
				Legendary: legendary,
			}
			return pokemon
		}
	}
	return (utils.PokemonData{})
}
