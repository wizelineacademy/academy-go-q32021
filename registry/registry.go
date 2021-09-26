package registry

import "github.com/nestorivan/academy-go-q32021/controller"


type registry struct{}

type Registry interface{
  NewAppController() controller.AppController
}

func NewRegistry() Registry{
  return &registry{}
}

func (r *registry) NewAppController() controller.AppController{
  return controller.AppController{
    Pokemon: r.NewPokemonController(),
  }
}