package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Marcxz/academy-go-q32021/usecase"
)

var (
	au = usecase.NewAddressUseCase()
)

type Address interface {
	ReadCSVAddress(http.ResponseWriter, *http.Request)
}

type c struct{}

func NewAddressController() Address {
	return &c{}
}

func (*c) ReadCSVAddress(w http.ResponseWriter, r *http.Request) {
	ad, err := au.ReadCSVAddress("")

	if err != nil {
		HandleError(w, r, err)
	}

	ja, err := json.Marshal(ad)

	if err != nil {
		HandleError(w, r, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ja)
}

func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	log.Fatalln(err)
}
