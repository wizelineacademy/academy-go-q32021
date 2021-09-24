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

// Adress - Interface for Address Controller
type Address interface {
	ReadCSVAddress(http.ResponseWriter, *http.Request)
}

type c struct{}

// NewAddressController - The constructor for a controller used at routes
func NewAddressController() Address {
	return &c{}
}

// ReadCSVAddress - Handler to read the all the Addresses from a csv file
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

// HandleError - Refactored func to report the errors in the controllers
func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	log.Fatalln(err)
}
