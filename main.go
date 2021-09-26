package main

import (
	"fmt"
	"mrobles_app/app"
)

func main() {
	app.RunApp()
	fmt.Println("Server listening on http://localhost:" + "5000")
}
