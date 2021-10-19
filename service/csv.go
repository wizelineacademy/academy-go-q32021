package service

import (
	"Project/config"
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetLeaderInCSV(country string) map[string]string {

	f, err := os.Open("./tmp/" + "result.csv")
	if err != nil {
		log.Error("Error reading ", err)
	}
	r := csv.NewReader(bufio.NewReader(f))
	i := 0

	for {
		record, err := r.Read()

		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if record[3] == country {

			result := map[string]string{
				"country": record[3],
				"name":    record[1],
				"rank":    record[0],
			}
			return result
		}
		i += 1
	}

	return map[string]string{"name": "Not found", "rank": "Not found"}
}

func ConcurrentRead(vars map[string]string) map[int][]string {

	items, err := strconv.Atoi(vars["items"])
	if err != nil {
		log.Error("There was an error converting items parameter to int: ", err)
	}

	itemsPerWorker, err := strconv.Atoi(vars["itemsPerWorker"])
	if err != nil {
		log.Error("There was an error converting items parameter to int: ", err)
	}

	jobs := make(chan []string, items)
	results := make(chan []string, items)

	for i := 0; i <= (items / itemsPerWorker); i++ {
		go Worker(i, jobs, results, vars["type"])
	}

	csvFile, err := os.Open(config.CsvPath)
	if err != nil {
		log.Error("Error opening CSV: ", err)
	}
	log.Debug("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Error("Error reading CSV: ", err)
	}

	for index, line := range csvLines {

		jobs <- line

		if index >= items {
			break
		}
	}

	close(jobs)

	var result = make(map[int][]string)

	for a := 1; a <= items; a++ {

		t := <-results

		if t != nil {
			result[a] = t
		}
		log.Debug("Results: ", t)
	}

	log.Debug(result)

	return result
}

func Worker(id int, jobs <-chan []string, results chan<- []string, oddOrEven string) {
	for j := range jobs {
		log.Debug("Worker: ", id, " started  job", j)
		results <- TypeOfItem(j, oddOrEven)
		log.Debug("Worker: ", id, " finished job", j)
	}
}

func TypeOfItem(item []string, divisibility string) []string {

	if item[0] == "id" {
		return nil
	}

	if divisibility == "odd" {

		id, err := strconv.Atoi(item[0])
		if err != nil {
			log.Error("There was a problem converting string to int: ", err)
		}

		if id%2 == 1 {
			return item
		} else {
			return nil
		}

	} else if divisibility == "even" {

		id, err := strconv.Atoi(item[0])
		if err != nil {
			log.Error("There was a problem converting string to int: ", err)
		}

		if id%2 == 0 {
			return item
		} else {
			return nil
		}

	} else {
		log.Error("Bad type parameter")
		return nil
	}
}
