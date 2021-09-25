package common

import (
	"encoding/csv"
	"errors"
	"os"
)

// ReadCsv - Reads file from a given path, returns the slice of records
func ReadCsv(Filename string) ([][]string, error) {
	file, err := os.Open(Filename)
	if err != nil {
		return [][]string{}, errors.New("an error ocurred while opening the csv file")
	}
	defer file.Close()

	Records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return [][]string{}, errors.New("an error ocurred while reading the csv file")
	}
	return Records, nil
}
