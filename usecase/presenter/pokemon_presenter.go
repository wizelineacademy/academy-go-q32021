package presenter

import "github.com/macrojoe/academy-go-q32021/domain/model"

type PokemonPresenter interface {
	ResponsePokemons(u []*model.Pokemon) []*model.Pokemon
}
