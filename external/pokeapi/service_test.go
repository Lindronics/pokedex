//	Tests the pokeapi service.
//	I omitted the tests for the /pokemon-species endpoint and GetPokemonProfile, as they would look largely the same.
//	In a production environment they would of course be included.
package pokeapi

import (
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/lindronics/pokedex/external"
)

const (
	name = "pikachu"
	url  = "https://test.com"
	path = "/pokemon/" + name
)

func mockExternalCallResponse(t *testing.T, externalResponseCode int, externalResponseBody string) (*Pokemon, *external.CallError) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	os.Setenv(pokeApiUrlParam, url)
	defer os.Unsetenv(pokeApiUrlParam)

	httpmock.RegisterResponder("GET", url+path, httpmock.NewStringResponder(externalResponseCode, externalResponseBody))
	pokemon, err := getPokemon(name)

	callCountInfo := httpmock.GetCallCountInfo()
	if i := callCountInfo["GET "+url+path]; i != 1 {
		t.Errorf("Must call %s exactly once but called %d times", url+path, i)
	}

	return pokemon, err
}

// TestPokemonExternalCallSuccess tests error scenarios for unsuccessful calls to a mocked GET /pokemon external API
func TestPokemonExternalCallErrors(t *testing.T) {
	tables := []struct {
		externalResponseCode int
		externalResponseBody string
		expectedCode         int
	}{
		{404, "", 404},
		{400, "", 500},
		{403, "", 500},
		{500, "", 502},
		{504, "", 502},
		{200, `{"species": {"name" : "pikachu", "url": "https://asdf"}}`, 502},
		{200, `{"name": "pikachu", "species": {"name": "pikachu"}}`, 502},
		{200, "", 502},
		{200, "invalid json", 502},
	}
	for _, table := range tables {
		_, err := mockExternalCallResponse(t, table.externalResponseCode, table.externalResponseBody)
		if err == nil {
			t.Errorf("Must return error")
		} else if err.ResponseCode != table.expectedCode {
			t.Errorf("Must return %d but was %d", table.expectedCode, err.ResponseCode)
		}
	}
}

// TestPokemonExternalCallSuccess tests a successful call against a mocked GET /pokemon external API
func TestPokemonExternalCallSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockResponse := `{"name": "pikachu", "species": {"name" : "pikachu", "url": "https://asdf"}, "other-property": "other-value"}`

	pokemon, err := mockExternalCallResponse(t, 200, mockResponse)
	if err != nil {
		t.Errorf("Must not return error")
	}
	if pokemon == nil {
		t.Errorf("Response object must not be nil")
	}
}
