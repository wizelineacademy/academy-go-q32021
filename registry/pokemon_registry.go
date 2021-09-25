package registry

import (
	"github.com/macrojoe/academy-go-q32021/interface/controller"
	ip "github.com/macrojoe/academy-go-q32021/interface/presenter"
	ir "github.com/macrojoe/academy-go-q32021/interface/repository"
	"github.com/macrojoe/academy-go-q32021/usecase/interactor"
	up "github.com/macrojoe/academy-go-q32021/usecase/presenter"
	ur "github.com/macrojoe/academy-go-q32021/usecase/repository"
)

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository(), r.NewPokemonPresenter())
}

func (r *registry) NewPokemonRepository() ur.PokemonRepository {
	return ir.NewPokemonRepository()
}

func (r *registry) NewPokemonPresenter() up.PokemonPresenter {
	return ip.NewPokemonPresenter()
}
