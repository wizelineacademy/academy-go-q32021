package presenter

import (
	"github.com/nestorivan/academy-go-q32021/domain/model"
)


type PokemonPresenter interface {
  ResponsePokemon(pk []model.Pokemon) []model.Pokemon
}

type pokemonPresenter struct{}

func NewPokemonPresenter() PokemonPresenter {
  return &pokemonPresenter{}
}

func (pp *pokemonPresenter) ResponsePokemon(pkmn []model.Pokemon) []model.Pokemon {
  return pkmn
}