package service

import (
	"Project/config"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetUserID(userName string) string {

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

func GetUserLastGame(userSteamId string) *http.Response {

	playerLastMatchApi := config.AOE2ApiLastGame + userSteamId

	respLastGame, errLastGame := http.Get(playerLastMatchApi)
	if errLastGame != nil {
		log.Error("Error retrieving last match: ", errLastGame)
		return nil
	}

	return respLastGame
}

func GetLeaderboard(count string) map[string]interface{} {

	apiResponse, err := http.Get(config.AOE2ApiLeaderboard + count)
	if err != nil {
		log.Error("Error triying to consume Api:", err)
	}

	responseData, err := ioutil.ReadAll(apiResponse.Body)
	if err != nil {
		log.Error("Error reading leaderboard data:", err)
	}

	var result map[string]interface{}
	errUnmarshalId := json.Unmarshal(responseData, &result)

	if errUnmarshalId != nil {
		log.Error("Error triying to unmarshall user ID:", errUnmarshalId)
	}

	return result
}
