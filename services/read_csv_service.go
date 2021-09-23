package services

import (
	"academy-go-q32021/entities"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadCsv(id string) (entities.Pokemon, error) {
	// Open the file

	fmt.Println("reading file...")

	target, targetError := strconv.Atoi(id)
	if targetError != nil {
		fmt.Println("target error")
		return notFound(), errors.New("Pokemon id must be a valid number")
	}
	csvfile, err := os.Open("resources/pokemon.csv")
	if err != nil {
		fmt.Println("Couldn't open the csv file", err)
		return notFound(), errors.New("An error happened opening the CSV file")
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			return notFound(), nil
		}
		if err != nil {
			fmt.Println("An error has ocurred", err)
			return notFound(), errors.New("An unexpected error happened while processing the search")
		}

		rid, errId := strconv.Atoi(record[0])
		if errId == nil {
			if rid == target {
				return entities.Pokemon{Id: rid, Name: record[1]}, nil
			}
		} else {

			return notFound(), errors.New("An error happened reading an Id from the CSV file")
		}
	}
}

func notFound() entities.Pokemon {
	return entities.Pokemon{Id: 0, Name: "MissingNo"}
}
