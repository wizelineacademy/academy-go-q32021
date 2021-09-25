package repository

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strconv"

	"wizeline/model"
)

func GetUsersFromCSV(fileName string) (string, error){
	// open file
	f, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return "", errors.New("Error opening file " + fileName)
	}

	//close the file at the end of the program
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	//Read CSV file using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Println(err)
		return "", errors.New("Error reading file " + fileName)
	}

	userData, err := getUsers(data)
	if err != nil {
		log.Println(err)
		return "", errors.New("Error parsing data " + fileName)
	}
	jsonData, err := json.MarshalIndent(userData, "", "  ")
	if err != nil {
		log.Println(err)
		return "", errors.New("Error transforming csv to json " + fileName)
	}
	return string(jsonData), nil
}

func getUsers(data [][]string) ([]model.User, error) {
	// convert csv lines to array of structs
	var users []model.User
	for i, line := range data {
		if i > 0 { // omit header line
			var record model.User
			for j, field := range line {
				if j == 0 {
					var readerError error
					record.UserId,readerError = strconv.Atoi(field)
					if readerError != nil {
						return nil, errors.New("The UserId must be a numeric value ")
					}
				} else if j == 1 {
					record.Name = field
				} else if j == 2 {
					record.Email = field
				} else if j == 3 {
				record.Phone = field
			}
			}
			users = append(users, record)
		}
	}
	return users, nil
}

