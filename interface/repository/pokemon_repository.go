package repository

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

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
	records, err := csvReader.ReadAll()

	for _, record := range records {
		tmpID, _ := strconv.Atoi(record[0])
		pokemon := model.Pokemon{
			ID:   tmpID,
			Name: record[1],
			Type: record[2],
		}

		u = append(u, &pokemon)
	}

	if err != nil {
		return nil, err
	}

	return u, nil
}
