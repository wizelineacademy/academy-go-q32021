package services

import (
	"mrobles_app/common"
	"strconv"
)

func parseCsv() ([]common.Game, error) {
	records, err := common.ReadCsvFile("steam.csv")
	if err != nil {
		return make([]common.Game, 0), err
	}

	slice := make([]common.Game, len(records[0]))

	for _, rec := range records {
		id, _ := strconv.Atoi(rec[0])
		achievments, _ := strconv.Atoi(rec[5])
		price, _ := strconv.ParseFloat(rec[6], 64)

		game := common.Game{
			ID:           id,
			Price:        price,
			Name:         rec[1],
			ReleaseDate:  rec[2],
			Developer:    rec[3],
			Publisher:    rec[4],
			Achievements: achievments,
		}

		slice = append(slice, game)
	}

	return slice, nil
}

func FindGame(id int) (common.Game, error) {
	games, err := parseCsv()
	if err != nil {
		return common.Game{}, err
	}

	for _, g := range games {
		if g.ID == id {
			return g, nil
		}
	}

	return common.Game{}, nil
}

func FindGames() ([]common.Game, error) {
	games, err := parseCsv()
	return games[7:], err
}
