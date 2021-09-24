package models

type Route struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	From address `json:"from"`
	To   address `json:"to"`
	R    []point `json:"r"`
}
