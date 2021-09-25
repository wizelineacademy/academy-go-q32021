package repository

import "github.com/macrojoe/academy-go-q32021/domain/model"

type PokemonRepository interface {
	FindAll(u []*model.Pokemon) ([]*model.Pokemon, error)
}
