package registry

import (
	"github.com/nestorivan/academy-go-q32021/controller"
	interactor "github.com/nestorivan/academy-go-q32021/interactors"
	presenter "github.com/nestorivan/academy-go-q32021/presenters"
)

func (r *registry) NewPokemonController() controller.PokemonController {
  return controller.NewPokemonController(r.NewCsvInteractor())
}

func (r *registry) NewCsvInteractor() interactor.CsvInteractor {
  pp := r.NewPokemonPresenter()
  return interactor.NewCsvInteractor(pp)
}

func (r *registry) NewPokemonPresenter() presenter.PokemonPresenter {
  return presenter.NewPokemonPresenter()
}