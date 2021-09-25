package model

//Pokemons struct
//ID int unique identifier
//Name string pokemon name
//MainType string pokemon main type
//SecondType string pokemon second type.
type Pokemon struct {
	ID         int    `csv:"id"`
	Name       string `csv:"Name"`
	MainType   string `csv:"Type 1"`
	SecondType string `csv:"Type 2"`
}
