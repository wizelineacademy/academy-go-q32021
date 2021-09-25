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

func ReadCSV(w http.ResponseWriter, r *http.Request) {

	var resp = render.New()
	var Contestants []model.Contestant

	csvf, err := os.Open("../api/lmd.csv")
	if err != nil {
		fmt.Println(err)
		resp.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	csvLines, err := csv.NewReader(csvf).ReadAll()
	if err != nil {
		fmt.Println(err)
		resp.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
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
	resp.JSON(w, http.StatusOK, Contestants)
}
