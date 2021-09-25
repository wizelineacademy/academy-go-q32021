package common

import (
	"encoding/csv"
	"os"
)

func ReadCsvFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return make([][]string, 0), err
	}

	reader := csv.NewReader(file)
	return reader.ReadAll()
}
