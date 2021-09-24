package repositories

type Pokemon struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Type1        string `json:"type_1"`
	Type2        string `json:"type_2"`
	Total        int    `json:"total_points"`
	HP           int    `json:"hp"`
	Attack       int    `json:"attack"`
	Defense      int    `json:"defense"`
	SpeedAttack  int    `json:"speed_attack"`
	SpeedDefense int    `json:"speed_defense"`
	Speed        int    `json:"speed"`
	Generation   int    `json:"generation"`
	Legendary    string `json:"legendary"`
}
