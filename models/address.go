package models

type Address struct {
	Id int
	A  string `json:"a"`
	P  Point  `json:"p"`
}
