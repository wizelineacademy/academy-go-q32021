package services

import (
	"fmt"
	"strconv"

	"github.com/s1nuh3/academy-go-q32021/common"
	"github.com/s1nuh3/academy-go-q32021/models"
)

// GetUsersfromCSV - Returns a colection of model.users from a csv file
func GetUsersfromCSV() ([]models.Users, error) {
	rcd, err := common.ReadCsv("./repositories/files/usersdata.csv")
	if err != nil {
		return nil, err
	}

	var u []models.Users
	for _, r := range rcd {
		id, _ := strconv.Atoi(r[0])
		status, _ := strconv.ParseBool(r[4])
		data := models.Users{
			Id:     id,
			Name:   r[1],
			Email:  r[2],
			Gender: r[3],
			Status: status,
		}
		u = append(u, data)
	}
	fmt.Println(u)
	return u, nil
}

// GetUserbyIdfromCSV - Returns a user by id if it's found in a csv file
func GetUserbyIdfromCSV(id int) (models.Users, error) {
	u := models.Users{}
	rcd, err := common.ReadCsv("./repositories/files/emptyFile.csv")
	if err != nil {
		return u, err
	}

	for _, r := range rcd {
		i, _ := strconv.Atoi(r[0])
		if i == id {
			status, _ := strconv.ParseBool(r[4])
			u = models.Users{
				Id:     i,
				Name:   r[1],
				Email:  r[2],
				Gender: r[3],
				Status: status,
			}
		}
	}
	fmt.Println(u)
	return u, nil
}
