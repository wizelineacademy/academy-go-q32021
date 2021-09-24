package controller

import (
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{ "Message": "Hello, world!" }`)
	fmt.Println("Endpoint reached: helloWorld")
}
