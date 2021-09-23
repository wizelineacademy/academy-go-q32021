package repository

import (
	"fmt"
	"strings"

	infraestructure "github.com/Marcxz/academy-go-q32021/infraestructure/csv"
)

type csv interface {
	ReadCSVFile(f string) ([]string, error)
}

type cr struct{}

func NewCsvRepository() csv {
	return &cr{}
}

func (*cr) ReadCSVFile(f string) ([]string, error) {
	cl, err := infraestructure.ReadCSVFile(f)

	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	as := strings.Split(string(cl), "\n")

	return as, nil
}
