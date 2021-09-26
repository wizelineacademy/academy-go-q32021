package repository

import "github.com/nestorivan/academy-go-q32021/domain/model"


type PokemonRepository interface{
  FindAll(id string) ([]*model.Pokemon, error)
  InsertOne(*model.Pokemon) ([]*model.Pokemon, error)
}