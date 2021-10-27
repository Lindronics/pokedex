package pokeapi

const HabitatCave = "cave"

type Pokemon struct {
	Name    string      `json:"name" valid:"required"`
	Species Description `json:"species" valid:"required"`
}

type Description struct {
	Name string `json:"name" valid:"required"`
	Url  string `json:"url" valid:"required,url"`
}

type PokemonSpecies struct {
	Name        string       `json:"name" valid:"required"`
	Habitat     Description  `json:"habitat" valid:"required"`
	FlavorTexts []FlavorText `json:"flavor_text_entries" valid:"required"`
	IsLegendary bool         `json:"is_legendary"`
}

type FlavorText struct {
	Text     string      `json:"flavor_text" valid:"required"`
	Language Description `json:"language" valid:"required"`
	Version  Description `json:"version" valid:"required"`
}
