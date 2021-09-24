package interactor

import (
	"github.com/OrlandoRomo/academy-go-q32021/domain/model"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/presenter"
	"github.com/OrlandoRomo/academy-go-q32021/usercase/repository"
)

type listInteractor struct {
	ListRepository repository.ListRepository
	ListPresenter  presenter.ListPresenter
}

type ListInteractor interface {
	Get(term string) ([]*model.List, error)
}

func NewListInteractor(r repository.ListRepository, p presenter.ListPresenter) *listInteractor {
	return &listInteractor{r, p}
}
