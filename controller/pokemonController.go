package controller

import (
	"encoding/csv"
	"net/http"
	. "pokemon/domain/model"

	"github.com/gin-gonic/gin"
)

type PokemonController struct {
	*gin.Engine
}

func NewPokemonController(e *gin.Engine) *PokemonController {
	return &PokemonController{e}
}

func (this *PokemonController) GetPokemon() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pikachu = Pokemon{Id: "25", Name: "Pikachu"}
		ctx.JSON(200, pikachu)
	}
}

func (this *PokemonController) GetPokemonFromFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": http.StatusUnprocessableEntity,
				"error":  "Unable to get file",
			})
		}

		fileContent, _ := file.Open()
		r := csv.NewReader(fileContent) // <= Error message
		records, err := r.ReadAll()

		var result = make([]Pokemon, 0, 0)

		for _, record := range records {
			pkmn := Pokemon{
				Id:   record[0],
				Name: record[1],
			}
			result = append(result, pkmn)
		}
		c.JSON(200, result)
	}
}

func (this *PokemonController) Router() {
	this.Handle("GET", "/", this.GetPokemon())
	this.Handle("POST", "/", this.GetPokemonFromFile())
}
