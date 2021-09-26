package main

import (
	"github.com/nestorivan/academy-go-q32021/infrastructure/router"
	"github.com/nestorivan/academy-go-q32021/registry"
)


func main(){
  r:= registry.NewRegistry()

  rt := router.NewRouter(r.NewAppController())

  rt.Run(":8000")
}