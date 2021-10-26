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

	r.GET("/pokemon/:name", getBasicInformation)
	r.GET("/pokemon/translated/:name", getTranslatedInformation)

	r.Run()
}

func getBasicInformation(ctx *gin.Context) {
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

func getTranslatedInformation(ctx *gin.Context) {
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
	pokemon.Description, err = translation.TranslateText(pokemon.Description, translation.Shakespeare)
	if err != nil {
		ctx.IndentedJSON(err.ResponseCode, model.ErrorResponse{Message: err.Message})
		return
	}
	ctx.IndentedJSON(http.StatusOK, pokemon)
}
