package controllers

import "github.com/OrlandoRomo/academy-go-q32021/usercase/interactor"

type listController struct {
	listInteractor interactor.ListInteractor
}

type ListController interface {
	GetDefinitions() error
}

func (l *listController) GetDefinitions() error {
	return nil
}
