package model

type ListDefinition struct {
	List []List `json:"list"`
}

// Struct to unmarshal the Urban dictionary response of definitions based on a term
// API For more infomation: https://rapidapi.com/community/api/urban-dictionary/
type List struct {
	Definition  string   `json:"definition"`
	Permalink   string   `json:"permalink"`
	ThumbsUp    int64    `json:"thumbs_up"`
	SoundUrls   []string `json:"sound_urls"`
	Author      string   `json:"author"`
	Word        string   `json:"word"`
	Defid       int64    `json:"defid"`
	CurrentVote string   `json:"current_vote"`
	WrittenOn   string   `json:"written_on"`
	Example     string   `json:"example"`
	ThumbsDown  int64    `json:"thumbs_down"`
}
