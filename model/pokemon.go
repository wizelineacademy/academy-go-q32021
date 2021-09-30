package model

type PokemonData struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Type1     string `json:"type_1"`
	Type2     string `json:"type_2"`
	Legendary bool   `json:"legendary"`
}
