package app

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CSVReader() [][]string {

	// open the file
	csvFile, err := os.Open("../pokedex.csv")
	if err != nil {
		fmt.Println("error encountered opening csv file")
		panic(err.Error())
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Printf("error encountered opening csv file: %v", err.Error())
	}
	return csvLines
}
