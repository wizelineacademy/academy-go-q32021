package main

//Get gin
import (
	. "pokemon/controller"

	"github.com/gin-gonic/gin"
)

//Principal function
func main() {
	//R is the abbreviation of router
	r := gin.Default()
	//It's very simple here. It's very similar to the routing of DeNO and node
	NewPokemonController(r).Router()
	//Monitor port 8080
	r.Run(":8080")
}
