package controller

import (
	"net/http"

	"github.com/macrojoe/academy-go-q32021/domain/model"
	"github.com/macrojoe/academy-go-q32021/usecase/interactor"
)

type pokemonController struct {
	userInteractor interactor.PokemonInteractor
}

type PokemonController interface {
	GetPokemons(c Context) error
}

func NewPokemonController(us interactor.PokemonInteractor) PokemonController {
	return &pokemonController{us}
}

func (uc *pokemonController) GetPokemons(c Context) error {
	var u []*model.Pokemon

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
