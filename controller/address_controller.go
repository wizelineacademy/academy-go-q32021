package controller

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/Marcxz/academy-go-q32021/usecase"
	"github.com/gorilla/mux"
)

var (
	au = usecase.NewAddressUseCase();
)
type Address interface {
	ReadCSVAddress(s string)
}

type c struct{}

NewAddressController() Address {
	return &c{}
}

func (*c) ReadCSVAddress(w http.ResponseWriter, r *http.Request) {
	ad, err := au.ReadCSVAddress()

	if (err != nil) {
		HandleError(w, r, err)
	}
	
	ja, err := json.Marshal(ad)

	if (err != nil) {
		HandleError(w, r, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ja)

}

func HandleError(w httpResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	log.Fatalln(err)
}
