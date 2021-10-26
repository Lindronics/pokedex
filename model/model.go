package model

type PokemonResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Habitat     string `json:"habitat"`
	IsLegendary bool   `json:"is_legendary"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
