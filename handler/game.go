package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"

	"Project/config"
	"Project/model"
	"Project/response"
	"Project/service"
	"Project/utils"
)

func GetLeaderboardDefault(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	result := service.GetLeaderboard(config.ItemsDefault)

	response := response.NewResponseInterface(result, startTime)
	json.NewEncoder(w).Encode(response)
}

func GetLeaderboardCount(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	vars := mux.Vars(r)
	log.Debug("Received as params : ", vars["count"])

	result := service.GetLeaderboard(vars["count"])

	mapLeaderboard, err := json.Marshal(result["leaderboard"])
	if err != nil {
		panic(err)
	}

	var players []model.Player

	errUnmarshal := json.Unmarshal(mapLeaderboard, &players)

	if errUnmarshal != nil {
		log.Error("Error triying to unmarshall user ID:", errUnmarshal)
	}

	utils.CreateCSV(players)

	response := response.NewResponseInterface(result, startTime)
	json.NewEncoder(w).Encode(response)
}

func GetLeaderByCountry(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	vars := mux.Vars(r)
	log.Debug("Received as params : ", vars["country"])

	leader := service.GetLeaderInCSV(vars["country"])

	response := response.NewResponse(leader, startTime)
	json.NewEncoder(w).Encode(response)
}

func ReadItems(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	vars := mux.Vars(r)
	log.Debug("Received as params Type: ", vars["type"], ", Items: ", vars["items"], ", itemsPerWorker: ", vars["itemsPerWorker"])

	if vars["type"] == "odd" || vars["type"] == "even" {
		result := service.ConcurrentRead(vars)
		response := response.NewResponseIndex(result, startTime)
		json.NewEncoder(w).Encode(response)
	} else {
		response := response.ErrorResponse(0, "Type not allowed", startTime)
		json.NewEncoder(w).Encode(response)
	}

}
