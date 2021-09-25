package model

type Pokemon struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Type_2    string `json:"type_2"`
	Total     int    `json:"total"`
	HP        int    `json:"hp"`
	Attack    int    `json:"attack"`
	Defense   int    `json:"defense"`
	SpAttack  int    `json:"sp_attack"`
	SpDefense int    `json:"sp_defense"`
	Speed     int    `json:"speed"`
	Gen       int    `json:"generation"`
	Legendary bool   `json:"legendary"`
}
