package model

type ContestantDetailed struct {
	ID           int    `json:"ID"`
	Contestant   string `json:"Contestant"`
	RealName     string `json:"Real Name"`
	Age          int    `json:"Age"`
	CurrentCity  string `json:"Current City"`
	CurrentScore int    `json:"Score"`
	Bio          string `json:"Bio"`
}
