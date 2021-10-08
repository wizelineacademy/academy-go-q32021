package model

type Contestant struct {
	ID           int    `json:"ID"`
	Contestant   string `json:"Contestant"`
	RealName     string `json:"Real Name"`
	Age          int    `json:"Age"`
	CurrentCity  string `json:"Current City"`
	CurrentScore int    `json:"Score"`
}
