package model

type City struct {
	Id          int
	Name        string
	Countrycode string
}

func (City) TableName() string { return "cities" }
