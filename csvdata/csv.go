package csvdata

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Diegoplas/go-bootcamp-deliverable/config"
	"github.com/Diegoplas/go-bootcamp-deliverable/model"
)

func readCSVLines() ([][]string, error) {

	// open the file
	csvFile, err := os.Open(config.CSVPath)
	if err != nil {
		return nil, fmt.Errorf("error opening csv file: %v", err.Error())
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading csv file: %v", err.Error())
	}

	defer csvFile.Close()

	return csvLines, nil
}

func CheckIfPokemonAlreadyExists(index string) (model.PokemonData, error) {

	csvLines, err := readCSVLines()
	if err != nil {
		return model.PokemonData{}, fmt.Errorf("error converting string to int %v", err.Error())
	}

	for _, line := range csvLines {
		if line[0] != "" {

			pokemonID, err := strconv.Atoi(line[0])
			if err != nil {
				log.Println("Error converting string to int ")
				return model.PokemonData{}, fmt.Errorf("error with data %v", err.Error())
			}

			if line[0] == index {

				pokemon := model.PokemonData{
					ID:    pokemonID,
					Name:  line[1],
					Type1: line[2],
					Type2: line[3],
				}

				return pokemon, nil
			}
		}
	}

	return model.PokemonData{}, nil
}

func WritePokemonIntoCSV(externalPokemonData model.PokemonData) error {

	csvFile, err := os.Open(config.CSVPath)
	if err != nil {
		log.Printf("Error opening csv file: %s", err)
		return fmt.Errorf("CSV error")
	}

	writer := csv.NewWriter(csvFile)

	row := []string{

		strconv.Itoa(externalPokemonData.ID),
		strings.Title(externalPokemonData.Name),
		strconv.Itoa(externalPokemonData.Height),
		externalPokemonData.Type1,
		externalPokemonData.Type2,
	}

	errWriter := writer.Write(row)
	if errWriter != nil {
		log.Println("error writing record to file", errWriter)
		return fmt.Errorf("writter error")

	}

	writer.Flush()
	csvFile.Close()

	return nil
}
