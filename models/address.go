package models

//Address - struct for the address compose from id, an adress and a position (lat, lng)
type Address struct {
	Id int    `json:"id"`
	A  string `json:"a"`
	P  Point  `json:"p"`
}
