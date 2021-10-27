// Tests the API handlers. For the sake of my time, I omitted the test for GetPokemon
// and various error response tests. In a production environment, I would probably factor
// most of the test logic out, making it easier to quickly test more scenarios.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lindronics/pokedex/external"
	"github.com/lindronics/pokedex/external/translation"
	"github.com/lindronics/pokedex/model"
)

type MockPokeApi struct {
	responseObject *model.PokemonResponse
}

func (t *MockPokeApi) GetPokemonProfile(string) (*model.PokemonResponse, *external.CallError) {
	return t.responseObject, nil
}

type MockTranslator struct{}

func (t *MockTranslator) TranslateText(text string, translation translation.TranslatorType) (string, *external.CallError) {
	return text + " this translation just adds some text " + string(translation), nil
}

type ErrorMockTranslator struct{}

func (t *ErrorMockTranslator) TranslateText(text string, translation translation.TranslatorType) (string, *external.CallError) {
	return "", external.NewCallError(500, "some error", fmt.Errorf(""))
}

// TestGetTranslatedPokemon tests the GetTranslatedPokemon handler.
// Validates output based on various external API responses.
func TestGetTranslatedPokemon(t *testing.T) {
	tables := []struct {
		translator     translation.Provider
		responseObject *model.PokemonResponse
		description    string
	}{
		{
			&MockTranslator{},
			&model.PokemonResponse{
				Name:        "name",
				Description: "description",
				Habitat:     "cave",
				IsLegendary: false,
			},
			"description this translation just adds some text yoda",
		},
		{
			&MockTranslator{},
			&model.PokemonResponse{
				Name:        "name",
				Description: "description",
				Habitat:     "forest",
				IsLegendary: true,
			},
			"description this translation just adds some text yoda",
		},
		{
			&MockTranslator{},
			&model.PokemonResponse{
				Name:        "name",
				Description: "description",
				Habitat:     "forest",
				IsLegendary: false,
			},
			"description this translation just adds some text shakespeare",
		},
		{
			&ErrorMockTranslator{},
			&model.PokemonResponse{
				Name:        "name",
				Description: "description",
				Habitat:     "cave",
				IsLegendary: true,
			},
			"description",
		},
	}

	for _, table := range tables {
		ts := httptest.NewServer(setupServer(&MockPokeApi{table.responseObject}, table.translator))
		defer ts.Close()

		resp, err := http.Get(fmt.Sprintf("%s/pokemon/translated/%s", ts.URL, table.responseObject.Name))
		if err != nil {
			t.Errorf("Error during request")
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Invalid response")
		}
		var pokemon model.PokemonResponse
		err = json.Unmarshal(body, &pokemon)
		if err != nil {
			t.Errorf("Invalid response")
		}

		translatedResponseObject := *table.responseObject
		translatedResponseObject.Description = table.description

		if pokemon != translatedResponseObject {
			t.Errorf("Expected %v but got %v", translatedResponseObject, pokemon)
		}
	}
}
