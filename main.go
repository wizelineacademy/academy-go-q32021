package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/gorilla/mux"
)

type Pokemon struct {
	Id         int    `csv:"id"`
	Name       string `csv:"Name"`
	MainType   string `csv:"Type 1"`
	SecondType string `csv:"Type 2"`
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/pokemons", func(w http.ResponseWriter, r *http.Request) {
		pokemonsFile, err := os.OpenFile("./utils/pokemon.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		defer pokemonsFile.Close()
		pokemons := []*Pokemon{}

		if err := gocsv.UnmarshalFile(pokemonsFile, &pokemons); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(pokemons)
	})
	myRouter.HandleFunc("/pokemons/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		pokemonsFile, err := os.Open("./utils/pokemon.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer pokemonsFile.Close()
		pokemon := []*Pokemon{}
		reader := csv.NewReader(pokemonsFile)
		var headers []string
		var pokemonCsvString string
		i := 0
		for {
			record, err := reader.Read()
			if i == 0 {
				headers = record
				i++
				pokemonCsvString = strings.Join(headers, ",") + "\n"
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			if record[0] == id {
				pokemonCsvString += strings.Join(record, ",")
				fmt.Printf("%v %T \n", record, record)
				fmt.Printf("%v %T \n", pokemonCsvString, pokemonCsvString)
				break
			}
		}
		if err := gocsv.UnmarshalString(pokemonCsvString, &pokemon); err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(pokemon)
	})
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
