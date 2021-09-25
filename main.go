package main

import (
	"log"
	"net/http"
	"os"

	"github.com/s1nuh3/academy-go-q32021/app"
)

func main() {
	app := app.New()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("App running..")
	err := http.ListenAndServe(":8889", nil)
	errHandler(err)
}

func errHandler(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
