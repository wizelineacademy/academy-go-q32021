package main

import (
	"fmt"

	"github.com/macrojoe/academy-go-q32021/infrastructure/router"
)

func Hola() {

}
func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	router.HandleRequests()
}
