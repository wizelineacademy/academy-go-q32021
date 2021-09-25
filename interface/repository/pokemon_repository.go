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
	first := true
	for _, record := range records {
		if first {
			first = false
			continue
		}
		tmpID, _ := strconv.Atoi(record[0])
		tmpTotal, _ := strconv.Atoi(record[4])
		tmpHp, _ := strconv.Atoi(record[5])
		tmpAttack, _ := strconv.Atoi(record[6])
		tmpDefense, _ := strconv.Atoi(record[7])
		tmpSpAttack, _ := strconv.Atoi(record[8])
		tmpSpDefense, _ := strconv.Atoi(record[9])
		tmpSpeed, _ := strconv.Atoi(record[10])
		tmpGeneration, _ := strconv.Atoi(record[11])
		tmpLegendary, _ := strconv.ParseBool(record[12])

		pokemon := model.Pokemon{
			ID:        tmpID,
			Name:      record[1],
			Type:      record[2],
			Type_2:    record[3],
			Total:     tmpTotal,
			HP:        tmpHp,
			Attack:    tmpAttack,
			Defense:   tmpDefense,
			SpAttack:  tmpSpAttack,
			SpDefense: tmpSpDefense,
			Speed:     tmpSpeed,
			Gen:       tmpGeneration,
			Legendary: tmpLegendary,
		}

		u = append(u, &pokemon)
	}

	if err != nil {
		return nil, err
	}

	return u, nil
}
