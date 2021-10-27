package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lindronics/pokedex/external/pokeapi"
	"github.com/lindronics/pokedex/external/translation"
	"github.com/lindronics/pokedex/model"
)

func main() {
	setupServer(pokeapi.NewHttpProvider(), translation.NewHttpProvider()).Run()
}

func setupServer(pokeapiProvider pokeapi.Provider, translator translation.Provider) *gin.Engine {
	r := gin.Default()
	r.GET("/pokemon/:name", getPokemon(pokeapiProvider))
	r.GET("/pokemon/translated/:name", getTranslatedPokemon(pokeapiProvider, translator))
	return r
}

// getPokemon retrieves a basic Pokemon profile by calling the pokeapi external API
func getPokemon(pokeapiProvider pokeapi.Provider) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		name, ok := ctx.Params.Get("name")
		if !ok {
			ctx.IndentedJSON(400, model.ErrorResponse{Message: "Invalid request parameter"})
			return
		}
		pokemon, err := pokeapiProvider.GetPokemonProfile(name)
		if err != nil {
			ctx.IndentedJSON(err.ResponseCode, model.ErrorResponse{Message: err.Message})
			return
		}
		ctx.IndentedJSON(http.StatusOK, pokemon)
	}
	return fn
}

// getTranslatedPokemon retrieves a translated Pokemon profile by calling the pokeapi external API
// and the funtranslations external API.
// If the Pokemon's habitat is "cave" or it is legendary, the translation will be "Yoda", else "Shakespeare"
func getTranslatedPokemon(pokeapiProvider pokeapi.Provider, translator translation.Provider) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		name, ok := ctx.Params.Get("name")
		if !ok {
			ctx.IndentedJSON(400, model.ErrorResponse{Message: "Invalid request parameter"})
			return
		}
		pokemon, err := pokeapiProvider.GetPokemonProfile(name)
		if err != nil {
			ctx.IndentedJSON(err.ResponseCode, model.ErrorResponse{Message: err.Message})
			return
		}

		var translatorType translation.TranslatorType
		if pokemon.Habitat == pokeapi.HabitatCave || pokemon.IsLegendary {
			translatorType = translation.Yoda
		} else {
			translatorType = translation.Shakespeare
		}

		translatedDescription, err := translator.TranslateText(pokemon.Description, translatorType)
		if err == nil && translatedDescription != "" {
			pokemon.Description = translatedDescription
		}
		ctx.IndentedJSON(http.StatusOK, pokemon)
	}
	return fn
}
