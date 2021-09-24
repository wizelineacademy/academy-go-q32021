package usecase

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Marcxz/academy-go-q32021/models"
	"github.com/Marcxz/academy-go-q32021/repository"
)

type Address interface {
	readCSVAddress(string) ([]models.Address, error)
}

var (
	cr = repository.NewCsvRepository()
	as = make([]models.Address, 0)
)

type auc struct{}

func NewAddressUseCase() Address {
	return &auc{}
}

func validate(i int, l string) error {
	al := strings.Split(l, "|")
	if len(al) != 4 {
		return fmt.Errorf("the line at the index %d should be composed for 4 pipes", i)
	}

	_, err := strconv.Atoi(al[0])
	if err != nil {
		return fmt.Errorf("the id at the index %d should be integer", i)
	}

	_, err = strconv.ParseFloat(al[2], 64)
	if err != nil {
		return fmt.Errorf("the lat column at the index %s should be float", al[2])
	}

	_, err = strconv.ParseFloat(al[3], 64)
	if err != nil {
		return fmt.Errorf("the lng column at the index %s should be float", al[3])
	}

	return nil
}

/*
func (*auc) geoAddress(a string) (*models.Address, error) {
	ad := models.Address{
		A: a,
		P: models.Point{},
	}
	return &ad, errors.New("Winno")
}
*/

func (*auc) readCSVAddress(f string) ([]models.Address, error) {
	as = make([]models.Address, 0)
	
	if (f == "") {
		f = "address.csv"
	}

	cl, err := cr.readCSVFile(f)

	if err != nil {
		return nil, err
	}

	for i, l := range cl {
		err = validate(i, l)
		if err != nil {
			return nil, err
		}
		al := strings.Split(l, "|")
		id, _ := strconv.Atoi(al[0])
		an := al[1]
		lat, _ := strconv.ParseFloat(al[2], 64)
		lng, _ := strconv.ParseFloat(al[3], 64)
		a := &models.Address{
			Id: id,
			A:  an,
			P: models.Point{
				Lat: lat,
				Lng: lng,
			},
		}
		as = append(as, *a)
	}

	return as, nil
}
