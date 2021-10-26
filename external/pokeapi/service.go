package pokeapi

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/lindronics/pokedex/external"
	"github.com/lindronics/pokedex/model"
)

const (
	pokeApiUrl = "https://pokeapi.co/api/v2"
)

func GetPokemon(name string) model.FullPokemon {
	pokemon, _ := getPokemon(name)
	species, _ := getPokemonSpecies(pokemon.Species.Name)

	flavourTexts := make([]string, 0)
	for _, description := range species.FlavorTexts {
		if description.Language.Name == "en" {
			flavourTexts = append(flavourTexts, description.Text)
		}
	}
	return model.FullPokemon{
		Name:        pokemon.Name,
		Habitat:     species.Habitat.Name,
		IsLegendary: species.IsLegendary,
		Description: strings.Join(flavourTexts, " ")[:100], // TODO
	}
}

func getPokemon(name string) (*Pokemon, error) {
	body, err := external.GetCall(pokeApiUrl, "pokemon", name)
	if err != nil {
		return nil, err
	}

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		log.Fatal("Response body corrupted", err)
	}
	return &pokemon, nil
}

func getPokemonSpecies(name string) (*PokemonSpecies, error) {
	body, err := external.GetCall(pokeApiUrl, "pokemon-species", name)
	if err != nil {
		return nil, err
	}

	var species PokemonSpecies
	err = json.Unmarshal(body, &species)
	if err != nil {
		log.Fatal("Response body corrupted", err)
	}
	return &species, nil
}
