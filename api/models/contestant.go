package model

type Contestant struct {
	ID           int    `csv:"ID"`
	Contestant   string `csv:"Contestant"`
	RealName     string `csv:"Real Name"`
	Age          int    `csv:"Age"`
	CurrentCity  string `csv:"Current City"`
	CurrentScore int    `csv:"Score"`
}

func (Contestant) TableName() string {
	return "contestants"
}
