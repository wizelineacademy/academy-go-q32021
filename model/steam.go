package model

import "encoding/json"

type SteamIdData struct {
	steamid json.Number
	success json.Number
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
