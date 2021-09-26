package common

type Game struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	ReleaseDate  string  `json:"release_date"`
	Developer    string  `json:"developer"`
	Publisher    string  `json:"publisher"`
	Achievements int     `json:"achievements"`
	Price        float64 `json:"price"`
}
