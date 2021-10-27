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

const (
	pokeApiUrl       = "https://test.com"
	name             = "mewtwo"
	translatorApiUrl = "https://test2.com"
	description      = "some description"
	translationText  = " this translation just adds some text"
)

var responseObject = &model.PokemonResponse{
	Name:        name,
	Description: description,
	Habitat:     "cave",
	IsLegendary: true,
}

type MockPokeApi struct{}

func (t *MockPokeApi) GetPokemonProfile(string) (*model.PokemonResponse, *external.CallError) {
	return responseObject, nil
}

type MockTranslator struct{}

func (t *MockTranslator) TranslateText(text string, translation translation.Translator) (string, *external.CallError) {
	return text + " this translation just adds some text", nil
}

func TestGetPokemonSuccess(t *testing.T) {
	ts := httptest.NewServer(setupServer(&MockPokeApi{}, &MockTranslator{}))
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/pokemon/%s", ts.URL, name))
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

	if pokemon != *responseObject {
		t.Errorf("Expected %v but got %v", *responseObject, pokemon)
	}
}

func TestGetTranslatedPokemonSuccess(t *testing.T) {
	ts := httptest.NewServer(setupServer(&MockPokeApi{}, &MockTranslator{}))
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/pokemon/translated/%s", ts.URL, name))
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

	translatedResponseObject := *responseObject
	translatedResponseObject.Description = description + translationText

	if pokemon != translatedResponseObject {
		t.Errorf("Expected %v but got %v", responseObject, pokemon)
	}
}
