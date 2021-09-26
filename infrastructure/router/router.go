package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nestorivan/academy-go-q32021/controller"
)


func NewRouter(c controller.AppController) *gin.Engine {
  r := gin.Default()

  r.Handle("GET", "/pokemon/", c.Pokemon.GetPokemons())
  r.Handle("GET", "/pokemon/:id", c.Pokemon.GetPokemons())
  r.Handle("POST", "/pokemon", c.Pokemon.CreatePokemons())

  return r
}