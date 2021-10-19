package main

import (
	"Project/model"
	"Project/service"
	"encoding/json"
	"strconv"
	"testing"

	log "github.com/sirupsen/logrus"
)

// router.HandleFunc("/leaderboard", handler.GetLeaderboardDefault)
// 	router.HandleFunc("/leaderboard/{count}", handler.GetLeaderboardCount)

type leaderboardTest struct {
	arg1     string
	expected int
}

var leaderboardTests = []leaderboardTest{
	leaderboardTest{"1", 1},
	leaderboardTest{"1", 1},
	leaderboardTest{"2", 2},
	leaderboardTest{"20", 20},
	leaderboardTest{"100", 100},
	leaderboardTest{"6", 15},
	leaderboardTest{"3", 13},
	leaderboardTest{"abc", 0},
}

func TestLeaderboard(t *testing.T) {

	for _, test := range leaderboardTests {

		output := service.GetLeaderboard(test.arg1)

		mapLeaderboard, err := json.Marshal(output["leaderboard"])
		if err != nil {
			panic(err)
		}

		var players []model.Player

		errUnmarshal := json.Unmarshal(mapLeaderboard, &players)

		if errUnmarshal != nil {
			log.Error("Error triying to unmarshall user ID:", errUnmarshal)
		}

		if len(players) != test.expected {
			log.Errorf("Output %q not equal to expected %q", output, test.expected)

		}
	}

}

// 	router.HandleFunc("/leaderbycountry/{country}", handler.GetLeaderByCountry)
type leaderbycountryTest struct {
	arg1     string
	expected string
}

var leaderbycountryTests = []leaderbycountryTest{
	leaderbycountryTest{"MX", "MX"},
	leaderbycountryTest{"AR", "AR"},
	leaderbycountryTest{"CO", "CO"},
	leaderbycountryTest{"CA", "CA"},
	leaderbycountryTest{"CN", "CN"},
	leaderbycountryTest{"NO", "NO"},
	leaderbycountryTest{"3", ""},
	leaderbycountryTest{"abc", ""},
}

func TestLeaderbycountry(t *testing.T) {

	for _, test := range leaderbycountryTests {

		output := service.GetLeaderInCSV(test.arg1)

		if output["country"] != test.expected {
			log.Errorf("Output %q not equal to expected %q", output, test.expected)

		}
	}

}

// 	router.HandleFunc("/read/{type}/{items}/{itemsPerWorker}", handler.ReadItems)
type concurrentreadTest struct {
	arg1       string
	arg2, arg3 int
	expected   int
}

var concurrentreadTests = []concurrentreadTest{
	concurrentreadTest{"even", 100, 2, 100},
	concurrentreadTest{"odd", 100, 5, 100},
	concurrentreadTest{"even", 10, 2, 10},
	concurrentreadTest{"even123", 100, 2, 0},
	concurrentreadTest{"9809even", 100, 909090, 0},
}

func TestConcurrentread(t *testing.T) {

	for _, test := range concurrentreadTests {

		tMap := map[string]string{"type": test.arg1, "items": strconv.Itoa(test.arg2), "itemsPerWorker": strconv.Itoa(test.arg3)}
		output := service.ConcurrentRead(tMap)

		if len(output) != test.expected {
			log.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}

}
