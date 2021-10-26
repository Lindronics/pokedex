package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lindronics/pokedex/service"
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
	pokemon := service.GetFullInfo(name)
	ctx.JSON(http.StatusOK, pokemon)
}

func getTranslatedInformation(ctx *gin.Context) {
	name, ok := ctx.Params.Get("name")
	if !ok {
		log.Fatal("param")
	}
	pokemon := service.GetFullInfo(name)
	// TODO translation
	ctx.JSON(http.StatusOK, pokemon)
}
