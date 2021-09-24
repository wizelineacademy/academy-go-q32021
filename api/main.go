package main

import (
	"fmt"
	"log"
	"net/http"

	"api/router"
	"api/utils"
)

func main() {

	conf, err := utils.ReadConfig("config.json")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Server:", conf.Server.Address+"/api")

	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(conf.Server.Address, r))

}
