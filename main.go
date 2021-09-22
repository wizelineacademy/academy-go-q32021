package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Author struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var (
	authors []Author
)

const (
	servicePort = ":8000"
	csvFileName = "author.csv"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health", checkHealth)
	router.HandleFunc("/authors", getAuthors).Methods("GET")

	http.Handle("/", router)

	log.Println("Service running on port", servicePort)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func checkHealth(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)

	io.WriteString(resp, `{"status": "alive"}`)
}

func getAuthors(resp http.ResponseWriter, req *http.Request) {
	records, err := readData(csvFileName)
	if err != nil {
		log.Fatal(err)
	}

	authors = nil
	for _, record := range records {
		author := Author{
			ID:        record[0],
			FirstName: record[1],
			LastName:  record[2],
		}

		authors = append(authors, author)
	}

	resp.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(authors)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		io.WriteString(resp, `{"status": "error", "message": "Error marshalling authors"}`)
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func readData(csvFileName string) ([][]string, error) {
	f, err := os.Open(csvFileName)
	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
