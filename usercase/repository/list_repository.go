package repository

import "github.com/OrlandoRomo/academy-go-q32021/domain/model"

type ListRepository interface {
	FindDefinitionsList(l []*model.List) ([]*model.List, error)
}
