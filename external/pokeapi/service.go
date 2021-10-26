package pokeapi

import (
	"strings"

	"github.com/lindronics/pokedex/external"
	"github.com/lindronics/pokedex/model"
)

const (
	pokeApiUrl = "https://pokeapi.co/api/v2"
)

func GetPokemonProfile(name string) (*model.PokemonResponse, *external.HttpError) {
	pokemon, err := getPokemon(name)
	if err != nil {
		return nil, err
	}

	species, err := getPokemonSpecies(pokemon.Species.Name)
	if err != nil {
		return nil, err
	}

	flavourTexts := make([]string, 0)
	for _, description := range species.FlavorTexts {
		if description.Language.Name == "en" {
			flavourTexts = append(flavourTexts, description.Text)
		}
	}

	return &model.PokemonResponse{
		Name:        pokemon.Name,
		Habitat:     species.Habitat.Name,
		IsLegendary: species.IsLegendary,
		Description: strings.Join(flavourTexts, " ")[:100], // TODO
	}, nil
}

func getPokemon(name string) (*Pokemon, *external.HttpError) {
	var pokemon Pokemon
	err := external.GetCall(pokeApiUrl, "pokemon", name, &pokemon)
	return &pokemon, err
}

func getPokemonSpecies(name string) (*PokemonSpecies, *external.HttpError) {
	var species PokemonSpecies
	err := external.GetCall(pokeApiUrl, "pokemon-species", name, &species)
	return &species, err
}
