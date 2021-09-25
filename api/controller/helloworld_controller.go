package controller

import (
	"fmt"
	"net/http"

	"github.com/unrolled/render"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	var resp = render.New()

	fmt.Println("Endpoint reached: helloWorld")
	resp.JSON(w, http.StatusBadRequest, map[string]string{"message": "hello, world!"})
}
