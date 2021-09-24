package infraestructure

import (
	"fmt"
	"io/ioutil"
)

// ReadCSVFile - Read a CSV file with a filename specified
func ReadCSVFile(f string) ([]byte, error) {
	p := fmt.Sprintf("files/%s", f)
	l, err := ioutil.ReadFile(p)

	if err != nil {
		return nil, err
	}

	return l, nil
}
