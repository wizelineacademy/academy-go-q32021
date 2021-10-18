package handler

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"

	"Project/config"
	"Project/model"
	"Project/response"
	"Project/utils"
)

func GetLastGameData(w http.ResponseWriter, r *http.Request) {

	start_time := time.Now()

	vars := mux.Vars(r)
	log.Debug("Received as params : ", vars["user"])

	userSteamId := GetUserID(vars["user"])
	if userSteamId == "error" {
		response := response.ErrorResponse(http.StatusInternalServerError, "Error with user name", start_time)
		json.NewEncoder(w).Encode(response)
		return
	}

	playerLastMatchApi := config.AOE2ApiLastGame + userSteamId

	respLastGame, errLastGame := http.Get(playerLastMatchApi)
	if errLastGame != nil {
		log.Error("Error retrieving last match: ", errLastGame)
		response := response.ErrorResponse(http.StatusInternalServerError, "There was an error", start_time)
		json.NewEncoder(w).Encode(response)
		return
	}

	//We Read the response body on the line below.
	bodyRespLastGame, errLastGame := ioutil.ReadAll(respLastGame.Body)
	if errLastGame != nil {
		log.Error("Error reading last match : ", errLastGame)
		response := response.ErrorResponse(http.StatusInternalServerError, "There was an error", start_time)
		json.NewEncoder(w).Encode(response)
		return
	}
	log.Debug(string(bodyRespLastGame))

	var resultLastGame map[string]interface{}
	lastGameData := json.Unmarshal(bodyRespLastGame, &resultLastGame)
	if lastGameData != nil {
		log.Error(lastGameData)
	}

	response := response.NewResponseInterface(resultLastGame, start_time)
	json.NewEncoder(w).Encode(response)
}

func GetUserID(userName string) string {
	// userName := "Zotarix1"

	customApi := config.SteamApiURL + "?key=" + config.SteamApiKey + "&vanityurl=" + userName

	respPlayerID, errPlayerID := http.Get(customApi)
	if errPlayerID != nil {
		log.Error("Error triying to get user ID:", errPlayerID)
		return "error"
	}

	//We Read the response body on the line below.
	bodyPlayerId, errBodyPlayerId := ioutil.ReadAll(respPlayerID.Body)
	if errBodyPlayerId != nil {
		log.Error("Error reading user data:", errBodyPlayerId)
		return "error"
	}

	var result map[string]map[string]json.Number
	errUnmarshalId := json.Unmarshal(bodyPlayerId, &result)

	if errUnmarshalId != nil {
		log.Error("Error triying to unmarshall user ID:", errUnmarshalId)
		return "error"
	}

	playerId := result["response"]["steamid"]

	return string(playerId)
}

func GetLeaderboard100(w http.ResponseWriter, r *http.Request) {

	start_time := time.Now()

	api := "https://aoe2.net/api/leaderboard?game=aoe2de&leaderboard_id=3&start=1&count=1"

	apiResponse, err := http.Get(api)
	if err != nil {
		log.Error("Error triying to consume Api:", err)
		// return "error"
	}

	responseData, err := ioutil.ReadAll(apiResponse.Body)
	if err != nil {
		log.Error("Error reading leaderboard data:", err)
		// return "error"
	}

	var result map[string]interface{}
	errUnmarshalId := json.Unmarshal(responseData, &result)

	if errUnmarshalId != nil {
		log.Error("Error triying to unmarshall user ID:", errUnmarshalId)
		// return "error"
	}

	response := response.NewResponseInterface(result, start_time)
	json.NewEncoder(w).Encode(response)
}

func GetLeaderboardCount(w http.ResponseWriter, r *http.Request) {

	start_time := time.Now()

	vars := mux.Vars(r)
	log.Debug("Received as params : ", vars["count"])

	api := "https://aoe2.net/api/leaderboard?game=aoe2de&leaderboard_id=3&start=1&count=" + vars["count"]

	apiResponse, err := http.Get(api)
	if err != nil {
		log.Error("Error triying to consume Api:", err)
		// return "error"
	}

	responseData, err := ioutil.ReadAll(apiResponse.Body)
	if err != nil {
		log.Error("Error reading leaderboard data:", err)
		// return "error"
	}

	var result map[string]interface{}

	errUnmarshalId := json.Unmarshal(responseData, &result)

	if errUnmarshalId != nil {
		log.Error("Error triying to unmarshall user ID:", errUnmarshalId)
		// return "error"
	}

	mapLeaderboard, err := json.Marshal(result["leaderboard"])
	if err != nil {
		panic(err)
	}

	var players []model.Player

	errUnmarshal := json.Unmarshal(mapLeaderboard, &players)

	if errUnmarshal != nil {
		log.Error("Error triying to unmarshall user ID:", errUnmarshal)
		// return "error"
	}

	log.Error(players[0].Name)

	utils.CreateCSV(players)

	response := response.NewResponseInterface(result, start_time)
	json.NewEncoder(w).Encode(response)
}

func GetLeaderByCountry(w http.ResponseWriter, r *http.Request) {

	start_time := time.Now()

	vars := mux.Vars(r)
	log.Debug("Received as params : ", vars["country"])

	leader := GetLeaderInCSV(vars["country"])

	log.Debug(leader)

	response := response.NewResponse(leader, start_time)
	json.NewEncoder(w).Encode(response)
}

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
		log.Error(record[3])
		if record[3] == country {

			result := map[string]string{
				"name": record[1],
				"rank": record[0],
			}
			return result
		}
		i += 1
	}

	return map[string]string{"name": "Not found", "rank": "Not found"}
}

func ReadItems(w http.ResponseWriter, r *http.Request) {
	start_time := time.Now()

	vars := mux.Vars(r)
	log.Error("Received as params Type: ", vars["type"], ", Items: ", vars["items"], ", items_per_workers: ", vars["items_per_workers"])

	result := ConcurrentRead(vars)

	response := response.NewResponseIndex(result, start_time)
	json.NewEncoder(w).Encode(response)
}

func ConcurrentRead(vars map[string]string) map[int][]string {

	items, err := strconv.Atoi(vars["items"])
	if err != nil {
		log.Error("There was an error converting items parameter to int: ", err)
	}

	itemsPerWorker, err := strconv.Atoi(vars["items_per_workers"])
	if err != nil {
		log.Error("There was an error converting items parameter to int: ", err)
	}

	jobs := make(chan []string, items)
	results := make(chan []string, items)

	for i := 0; i <= (items / itemsPerWorker); i++ {
		go Worker(i, jobs, results, vars["type"])
	}

	csvFile, err := os.Open("./tmp/result.csv")
	if err != nil {
		log.Println("Error opening CSV: ", err)
	}
	log.Println("Successfully Opened CSV file")
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
