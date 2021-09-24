package controller

import (
	model "api/models"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/unrolled/render"
)

var Resp = render.New()

func ReadCSV(w http.ResponseWriter, r *http.Request) {

	var Contestants []model.Contestant

	csvf, err := os.Open("../api/lmd.csv")
	if err != nil {
		fmt.Println(err)
	}

	csvLines, err := csv.NewReader(csvf).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	defer csvf.Close()

	for i, line := range csvLines {

		if i != 0 {

			id, _ := strconv.Atoi(line[0])
			age, _ := strconv.Atoi(line[3])
			score, _ := strconv.Atoi(line[5])

			cont := model.Contestant{
				ID:           id,
				Contestant:   line[1],
				RealName:     line[2],
				Age:          age,
				CurrentCity:  line[4],
				CurrentScore: score,
			}
			Contestants = append(Contestants, cont)
		}
	}

	fmt.Println("Endpoint reached: readCSV")
	Resp.JSON(w, http.StatusOK, Contestants)
}
