package router

import (
	"net/http"

	"Project/config"
	"Project/handler"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/leaderboard", handler.GetLeaderboardDefault)
	router.HandleFunc("/leaderboard/{count}", handler.GetLeaderboardCount)
	router.HandleFunc("/leaderbycountry/{country}", handler.GetLeaderByCountry)
	router.HandleFunc("/read/{type}/{items}/{itemsPerWorker}", handler.ReadItems)

	log.Error(http.ListenAndServe(config.Port, router))
}
