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
	pokemon := getPokemon(name)
	species := getPokemonSpecies(pokemon.Species.Name)

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
		Description: strings.Join(flavourTexts, " ")[:100],
	}
}

func getPokemon(name string) Pokemon {
	body := external.GetCall("pokemon", name)

	var pokemon Pokemon
	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		log.Fatal("Response body corrupted")
	}
	return pokemon
}

func getPokemonSpecies(name string) PokemonSpecies {
	body := external.GetCall("pokemon-species", name)

	var species PokemonSpecies
	err := json.Unmarshal(body, &species)
	if err != nil {
		log.Fatal("Response body corrupted", err)
	}
	return species
}
