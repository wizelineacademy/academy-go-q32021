package router

import (
	"net/http"

	"Project/handler"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/civ/{id}/", handler.GetCivilizationById)
	router.HandleFunc("/lastGame/{user}", handler.GetLastGameData)
	router.HandleFunc("/leaderboard", handler.GetLeaderboard100)
	router.HandleFunc("/leaderboard/{count}", handler.GetLeaderboardCount)
	router.HandleFunc("/leaderbycountry/{country}", handler.GetLeaderByCountry)
	router.HandleFunc("/read/{type}/{items}/{items_per_workers}", handler.ReadItems)

	log.Error(http.ListenAndServe(":80", router))
}
