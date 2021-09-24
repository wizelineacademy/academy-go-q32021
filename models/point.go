package models

// Point - struct from a spatial point composed from a lat, lng coordinate
type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
