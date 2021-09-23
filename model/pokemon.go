package model

type Pokemon struct {
	Id         int    `csv:"id"`
	Name       string `csv:"Name"`
	MainType   string `csv:"Type 1"`
	SecondType string `csv:"Type 2"`
}
