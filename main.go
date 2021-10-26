package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lindronics/pokedex/external/pokeapi"
	"github.com/lindronics/pokedex/external/translation"
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
		log.Fatal("param")
	}
	pokemon := pokeapi.GetPokemon(name)
	ctx.JSON(http.StatusOK, pokemon)
}

func getTranslatedInformation(ctx *gin.Context) {
	name, ok := ctx.Params.Get("name")
	if !ok {
		log.Fatal("param")
	}
	pokemon := pokeapi.GetPokemon(name)
	pokemon.Description = translation.TranslateText(pokemon.Description, translation.Shakespeare)
	ctx.JSON(http.StatusOK, pokemon)
}
