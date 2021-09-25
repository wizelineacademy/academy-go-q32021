package presenter

import "github.com/macrojoe/academy-go-q32021/domain/model"

type pokemonPresenter struct {
}

type PokemonPresenter interface {
	ResponsePokemons(us []*model.Pokemon) []*model.Pokemon
}

func NewPokemonPresenter() PokemonPresenter {
	return &pokemonPresenter{}
}

func (up *pokemonPresenter) ResponsePokemons(us []*model.Pokemon) []*model.Pokemon {
	for _, u := range us {
		u.Name = "Pokemon: " + u.Name
	}
	return us
}
