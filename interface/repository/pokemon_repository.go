package repository

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/macrojoe/academy-go-q32021/domain/model"
)

type pokemonRepository struct {
}

type PokemonRepository interface {
	FindAll(u []*model.Pokemon) ([]*model.Pokemon, error)
}

func NewPokemonRepository() PokemonRepository {
	return &pokemonRepository{}
}

func (ur *pokemonRepository) FindAll(u []*model.Pokemon) ([]*model.Pokemon, error) {

	filePath := "csv/pokemon.csv"
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	u, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	if err != nil {
		return nil, err
	}

	return u, nil
}

func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
