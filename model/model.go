package model

type FullPokemon struct {
	Name        string `json:"name"`
	Description string
	Habitat     string
	IsLegendary bool
}
