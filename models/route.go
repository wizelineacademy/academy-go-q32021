package models

type Route struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	From Address `json:"from"`
	To   Address `json:"to"`
	R    []Point `json:"r"`
}
