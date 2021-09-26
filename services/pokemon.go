package services

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"

	"github.com/MoraAlex/academy-go-q32021/model"

	"github.com/gocarina/gocsv"
)

// return all pokemons in the csv
func GetAllPokemons() ([]*model.Pokemon, error) {
	pokemonsFile, err := os.OpenFile("./utils/pokemon.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer pokemonsFile.Close()
	pokemons := []*model.Pokemon{}

	if err := gocsv.UnmarshalFile(pokemonsFile, &pokemons); err != nil {
		log.Fatal(err)
	}
	return pokemons, nil
}

// return the pokemon that match with the id
func GetPokemonById(id string) (*model.Pokemon, error) {
	pokemonsFile, err := os.Open("./utils/pokemon.csv")
	if err != nil {
		return nil, err
	}
	defer pokemonsFile.Close()
	pokemon := []*model.Pokemon{}
	reader := csv.NewReader(pokemonsFile)
	var headers []string
	var pokemonCsvString string
	i := 0
	for {
		record, err := reader.Read()
		if i == 0 {
			headers = record
			i++
			pokemonCsvString = strings.Join(headers, ",") + "\n"
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] == id {
			pokemonCsvString += strings.Join(record, ",")
			break
		}
	}
	if err := gocsv.UnmarshalString(pokemonCsvString, &pokemon); err != nil {
		log.Fatalf("error Unmarshal: %v", err)
	}
	return pokemon[0], nil
}
