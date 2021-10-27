// Contains client service and models for the pokeapi external API
package pokeapi

import (
	"os"
	"strings"

	"github.com/lindronics/pokedex/external"
	"github.com/lindronics/pokedex/model"
)

const PokeApiUrlParam string = "POKEAPI_URL"

type Provider interface {
	GetPokemonProfile(string) (*model.PokemonResponse, *external.CallError)
}

type HttpProvider struct{}

// GetPokemonProfile retrieves a PokemonResponse object by calling /pokemon and /pokemon-species/
// If an error occurs, returns nil and an error object containing the status code to return.
func (p *HttpProvider) GetPokemonProfile(name string) (*model.PokemonResponse, *external.CallError) {
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
		if text.Language.Name == LanguageEn {
			flavourTexts = append(flavourTexts, text.Text)
		}
	}

	var description string
	if len(flavourTexts) == 0 {
		description = ""
	} else {
		description = strings.Replace(flavourTexts[0], "\n", " ", -1)
	}

	return &model.PokemonResponse{
		Name:        pokemon.Name,
		Habitat:     species.Habitat.Name,
		IsLegendary: species.IsLegendary,
		Description: description,
	}, nil
}

// getPokemon calls /pokemon and returns a Pokemon or an error
func getPokemon(name string) (*Pokemon, *external.CallError) {
	var pokemon Pokemon
	err := external.GetCall(os.Getenv(PokeApiUrlParam), "pokemon", name, &pokemon)
	return &pokemon, err
}

// getPokemon calls /pokemon-species and returns a PokemonSpecies or an error
func getPokemonSpecies(name string) (*PokemonSpecies, *external.CallError) {
	var species PokemonSpecies
	err := external.GetCall(os.Getenv(PokeApiUrlParam), "pokemon-species", name, &species)
	return &species, err
}
