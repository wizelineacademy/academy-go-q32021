package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nestorivan/academy-go-q32021/domain/model"
	interactor "github.com/nestorivan/academy-go-q32021/interactors"
)


type pokemonController struct {
  csvInteractor interactor.CsvInteractor
}

type PokemonController interface {
  GetPokemons() gin.HandlerFunc
  CreatePokemons() gin.HandlerFunc
}

func NewPokemonController(ci interactor.CsvInteractor) PokemonController{
  return &pokemonController{ci}
}

func (pk *pokemonController) GetPokemons() gin.HandlerFunc {
  return func(c *gin.Context) {
    id := c.Param("id")
    pkml, err := pk.csvInteractor.ReadCsv()

    if (err != nil){
      c.AbortWithStatus(http.StatusInternalServerError)
    }

    if (id == ""){
      c.JSON(http.StatusOK, pkml)
      return
    }

    pkm := model.Pokemon{};

    for _, p := range pkml{
      if (p.Id == id){
        pkm = p
      }
    }

    c.JSON(http.StatusOK, pkm)
  }
}

func (pk *pokemonController) CreatePokemons() gin.HandlerFunc {
  return func(c *gin.Context) {
    var pkmn model.Pokemon

    err := c.Bind(&pkmn)

    if (err != nil){
      c.Status(http.StatusBadRequest)
    }

    pk.csvInteractor.WriteCsv(pkmn)

    c.Status(http.StatusOK)
  }
}