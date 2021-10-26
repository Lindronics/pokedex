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
	for _, text := range species.FlavorTexts {
		if text.Language.Name == "en" {
			flavourTexts = append(flavourTexts, text.Text)
		}
	}
	description := strings.Replace(strings.Join(flavourTexts, " "), "\n", " ", -1)

	return &model.PokemonResponse{
		Name:        pokemon.Name,
		Habitat:     species.Habitat.Name,
		IsLegendary: species.IsLegendary,
		Description: description,
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
