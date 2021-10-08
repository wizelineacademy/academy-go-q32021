package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/emamex98/academy-go-q32021/model"
	"github.com/emamex98/academy-go-q32021/utils"
	"github.com/gorilla/mux"

	"github.com/unrolled/render"
)

func GetContestans(w http.ResponseWriter, r *http.Request) {

	resp := render.New()

	contestants, err := fetchContestansFromCSV(resp, w)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Endpoint reached: /contestants")
	resp.JSON(w, http.StatusOK, contestants)
}

func GetSingleContestant(w http.ResponseWriter, r *http.Request) {

	resp := render.New()
	args := mux.Vars(r)

	id, err := strconv.Atoi(args["id"])
	if err != nil {
		returnError400(resp, w, err)
		return
	}

	contestants, err := fetchContestansFromCSV(resp, w)
	if err != nil {
		fmt.Println(err)
	}

	for i := range contestants {
		if contestants[i].ID == id {
			resp.JSON(w, http.StatusOK, contestants[i])
			return
		}
	}

	fmt.Println("Endpoint reached: /contestants/" + strconv.Itoa(id))
	resp.JSON(w, http.StatusNotFound, map[string]string{"error": "id not found"})
}

func fetchContestansFromCSV(resp *render.Render, w http.ResponseWriter) ([]model.Contestant, error) {

	var Contestants []model.Contestant

	csvLines, err := utils.ReadCSV("../api/lmd.csv")
	if err != nil {
		returnError500(resp, w, err)
		return nil, err
	}

	for i, line := range csvLines {

		if i == 0 {
			continue
		}

		id, err := strconv.Atoi(line[0])
		if err != nil {
			returnError400(resp, w, err)
			return nil, err
		}

		age, err := strconv.Atoi(line[3])
		if err != nil {
			returnError400(resp, w, err)
			return nil, err
		}

		score, err := strconv.Atoi(line[5])
		if err != nil {
			returnError400(resp, w, err)
			return nil, err
		}

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

	return Contestants, nil
}

func returnError400(resp *render.Render, w http.ResponseWriter, err error) {
	fmt.Println(err)
	resp.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
}

func returnError500(resp *render.Render, w http.ResponseWriter, err error) {
	fmt.Println(err)
	resp.JSON(
		w,
		http.StatusInternalServerError,
		map[string]string{
			"error": "Something happened while processing your request, try again"})
}
