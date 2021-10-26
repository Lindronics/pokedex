package pokeapi

type Pokemon struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Species     Description `json:"species"`
}

type Description struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonSpecies struct {
	Name        string       `json:"name"`
	Habitat     Description  `json:"habitat"`
	FlavorTexts []FlavorText `json:"flavor_text_entries"`
	IsLegendary bool         `json:"is_legendary"`
}

type FlavorText struct {
	Text     string      `json:"flavor_text"`
	Language Description `json:"language"`
	Version  Description `json:"version"`
}
