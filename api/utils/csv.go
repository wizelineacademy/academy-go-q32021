package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSV(path string) ([][]string, error) {

	csvf, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	csvLines, err := csv.NewReader(csvf).ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer csvf.Close()
	return csvLines, nil
}
