package utils

import (
	"Project/model"
	"encoding/csv"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func CreateCSV(data []model.Player) {

	file, err := os.Create("./tmp/result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"id", "name", "clan", "country", "games", "losses", "wins"}
	writer.Write(headers)

	for _, value := range data {

		t := make([]string, 0, 6)

		t = append(t, strconv.Itoa(value.Rank))
		t = append(t, value.Name)
		t = append(t, value.Clan)
		t = append(t, value.Country)
		t = append(t, strconv.Itoa(value.Games))
		t = append(t, strconv.Itoa(value.Losses))
		t = append(t, strconv.Itoa(value.Wins))

		_ = writer.Write(t)
		checkError("Cannot write to file", err)
	}

}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
