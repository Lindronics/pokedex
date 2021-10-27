package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lindronics/pokedex/external/pokeapi"
	"github.com/lindronics/pokedex/external/translation"
	"github.com/lindronics/pokedex/model"
)

func main() {
	r := gin.Default()

	r.GET("/pokemon/:name", getPokemon)
	r.GET("/pokemon/translated/:name", getTranslatedPokemon)

	r.Run()
}

// getPokemon retrieves a basic Pokemon profile by calling the pokeapi external API
func getPokemon(ctx *gin.Context) {
	name, ok := ctx.Params.Get("name")
	if !ok {
		ctx.IndentedJSON(400, model.ErrorResponse{Message: "Invalid request parameter"})
		return
	}
	pokemon, err := pokeapi.GetPokemonProfile(name)
	if err != nil {
		ctx.IndentedJSON(err.ResponseCode, model.ErrorResponse{Message: err.Message})
		return
	}
	ctx.IndentedJSON(http.StatusOK, pokemon)
}

// getTranslatedPokemon retrieves a translated Pokemon profile by calling the pokeapi external API
// and the funtranslations external API.
// If the Pokemon's habitat is "cave", the translation will be "Yoda", else "Shakespeare"
func getTranslatedPokemon(ctx *gin.Context) {
	name, ok := ctx.Params.Get("name")
	if !ok {
		ctx.IndentedJSON(400, model.ErrorResponse{Message: "Invalid request parameter"})
		return
	}
	pokemon, err := pokeapi.GetPokemonProfile(name)
	if err != nil {
		ctx.IndentedJSON(err.ResponseCode, model.ErrorResponse{Message: err.Message})
		return
	}
	switch pokemon.Habitat {
	case pokeapi.HabitatCave:
		pokemon.Description, err = translation.TranslateText(pokemon.Description, translation.Yoda)
	default:
		pokemon.Description, err = translation.TranslateText(pokemon.Description, translation.Shakespeare)
	}
	if err != nil {
		ctx.IndentedJSON(err.ResponseCode, model.ErrorResponse{Message: err.Message})
		return
	}
	ctx.IndentedJSON(http.StatusOK, pokemon)
}
