package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gocarina/gocsv"
)

type Pokemon struct {
	Id         int    `csv:"id"`
	Name       string `csv:"Name"`
	MainType   string `csv:"Type 1"`
	SecondType string `csv:"Type 2"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, peticion *http.Request) {
		pokemonsFile, err := os.OpenFile("./pokemon.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
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
	direccion := ":8080"
	fmt.Println("Servidor listo escuchando en " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))
}
