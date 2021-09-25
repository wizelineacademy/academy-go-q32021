package registry

import (
	"github.com/macrojoe/academy-go-q32021/interface/controller"
)

type registry struct {
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewPokemonController()
}
