package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/emamex98/academy-go-q32021/model"
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

func WriteChangesToCSV(path string, records []model.Contestant) error {

	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	header := strings.Split("ID,Contestant,Real Name,Age,Current City,Score,Bio", ",")
	if err := w.Write(header); err != nil {
		fmt.Println(err)
		return err
	}

	for _, record := range records {

		row := []string{
			strconv.Itoa(record.ID),
			record.Contestant,
			record.RealName,
			strconv.Itoa(record.Age),
			record.CurrentCity,
			strconv.Itoa(record.CurrentScore),
			record.Bio}

		if err := w.Write(row); err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil

}
