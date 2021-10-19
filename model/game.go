package model

type LastGame struct {
	name       string
	country    string
	last_match LastGameData
}

type LastGameData struct {
	match_id    int
	num_players int
	map_type    int
}

type Player struct {
	Rank    int
	Country string
	Clan    string
	Name    string
	Games   int
	Wins    int
	Losses  int
}
