package model

type TypeInfo struct {
	Name string `json:"name"`
}

type Types struct {
	Type TypeInfo `json:"type"`
}

type PokemonExternalData struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Height int     `json:"height"`
	Types  []Types `json:"types"`
}
