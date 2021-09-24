package services

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/bifr0ns/academy-go-q32021/repositories"
)

func GetPokemonInfo(pokemonId string) (*repositories.Pokemon, int, error) {

	csvFile, err := os.Open("pokemon.csv")
	if err != nil {
		return nil, 500, err
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, 500, err
	}

	for _, line := range csvLines {
		id, _ := strconv.Atoi(line[0])
		total, _ := strconv.Atoi(line[4])
		hp, _ := strconv.Atoi(line[5])
		attack, _ := strconv.Atoi(line[6])
		defense, _ := strconv.Atoi(line[7])
		speedAttack, _ := strconv.Atoi(line[8])
		speedDefense, _ := strconv.Atoi(line[9])
		speed, _ := strconv.Atoi(line[10])
		generation, _ := strconv.Atoi(line[11])

		pokemon := repositories.Pokemon{
			Id:           id,
			Name:         line[1],
			Type1:        line[2],
			Type2:        line[3],
			Total:        total,
			HP:           hp,
			Attack:       attack,
			Defense:      defense,
			SpeedAttack:  speedAttack,
			SpeedDefense: speedDefense,
			Speed:        speed,
			Generation:   generation,
			Legendary:    line[12],
		}

		id, _ = strconv.Atoi(pokemonId)
		if pokemon.Id == id {

			return &pokemon, 200, nil
		}
	}

	return nil, 404, errors.New("pokemon not found")
}
