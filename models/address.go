package models

type address struct {
	Id int    `json:"id"`
	A  string `json:"a"`
	P  point  `json:"p"`
}
