package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emamex98/academy-go-q32021/config"
	"github.com/emamex98/academy-go-q32021/router"
)

func main() {

	conf, err := config.ReadConfig("config.json")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Server:", conf.Server.Address+"/api")

	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(conf.Server.Address, r))

}
