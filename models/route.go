package models

// Route - struct from a route. The route has an id, name of the route, from point, to point and a collection of points that make the route
type Route struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	From Address `json:"from"`
	To   Address `json:"to"`
	R    []Point `json:"r"`
}
