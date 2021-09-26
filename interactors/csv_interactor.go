package interactor

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/nestorivan/academy-go-q32021/domain/model"
	presenter "github.com/nestorivan/academy-go-q32021/presenters"
)

var fileName string = "pokemon.csv"

type csvInteractor struct {
  PokemonPresenter presenter.PokemonPresenter
}

type CsvInteractor interface {
  ReadCsv()([]model.Pokemon, error)
  WriteCsv(pkmn model.Pokemon) ([]model.Pokemon, error)
}

func NewCsvInteractor(pp presenter.PokemonPresenter) CsvInteractor {
  return &csvInteractor{pp}
}

func getCsvValues (fn string) ([]model.Pokemon, error) {
  file,err := os.Open(fn)


  if err != nil{
    return nil, err
  }

  r := csv.NewReader(file)

  values,err := r.ReadAll()

  if err != nil {
    return nil, err
  }

  pkmList := []model.Pokemon{}

  for _, p := range values {
    pm := model.Pokemon{
      Id: p[0],
      Name: p[1],
    }
    pkmList = append(pkmList, pm)
  }
  defer file.Close()

  return pkmList, nil
}

func(ci *csvInteractor) saveCsvValues (fn string, pkm []model.Pokemon) error {
  // file,err := os.Open(fn)

  file, err := os.Open("pokemon.csv")

  if (err != nil){
    fmt.Println(err)
    return err
  }

  w := csv.NewWriter(file)

  for _, p := range pkm {
    var row []string
    row = append(row, p.Id)
    row = append(row, p.Name)
    w.Write(row)
  }

  w.Flush()

  defer func(){
    file.Close()
  }()

  return nil
}

func (ci *csvInteractor) ReadCsv() ([]model.Pokemon, error) {
  pkmList, err := getCsvValues(fileName)

  if err != nil{
    return nil, err
  }

  return ci.PokemonPresenter.ResponsePokemon(pkmList), nil
}

func (ci *csvInteractor) WriteCsv(pkmn model.Pokemon) ([]model.Pokemon, error ) {
  pkmList, err := getCsvValues(fileName)

  if err != nil {
    return nil, err
  }

  pkmList = append(pkmList, pkmn);

  e := ci.saveCsvValues(fileName, pkmList)

  if e != nil{
    return nil, err
  }

  return ci.PokemonPresenter.ResponsePokemon(pkmList), nil
}