package models

type Address struct {
	Id int `json:"id"`
	A  string `json:"a"`
	P  Point  `json:"p"`
}
