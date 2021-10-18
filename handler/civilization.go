package handler

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"

	"Project/config"
	"Project/response"
	"Project/utils"
)

func GetCivilizationById(w http.ResponseWriter, r *http.Request) {

	start_time := time.Now()

	vars := mux.Vars(r)
	log.Debug("Received as params : ", vars["id"])

	intVar, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := response.ErrorResponse(http.StatusInternalServerError, "Civilization ID not supported", start_time)
		json.NewEncoder(w).Encode(response)
		return
	}

	fileUrl := config.CsvURL
	fileErr := utils.DownloadFile("./tmp/"+config.CsvFileName, config.CsvURL)
	if fileErr != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// log.Error("Error Downloading file", err)

		response := response.ErrorResponse(http.StatusInternalServerError, "Error", start_time)
		json.NewEncoder(w).Encode(response)
		return
	}

	log.Debug("Downloaded: " + fileUrl)
	info := readCivData(intVar)

	log.Debug("NUM: ", intVar)
	log.Debug("Civ result: ", info)

	response := response.NewResponse(info, start_time)
	json.NewEncoder(w).Encode(response)
	return
}

func readCivData(id int) map[string]string {

	f, err := os.Open("./tmp/" + config.CsvFileName)
	if err != nil {
		log.Error("Error reading ", err)
	}
	r := csv.NewReader(bufio.NewReader(f))
	i := 0
	m := make(map[string]string)

	for {
		record, err := r.Read()

		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if i == id {
			m["id"] = record[0]
			m["name"] = record[1]
			m["specialty"] = record[3]
			m["uniqueUnit"] = record[4]
			break
		}

		i += 1
	}

	return m
}
