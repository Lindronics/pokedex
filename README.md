# Pokedex API exercise

[![Docker Image CI](https://github.com/Lindronics/pokedex/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/Lindronics/pokedex/actions/workflows/build.yml)

Contains a Pokedex REST API.

## How to run

If you have Docker, execute:

```$ docker-compose up --build```

Otherwise, you can run it by executing:

```$ go run .```

Note that if you aren't using Docker you'll have to set the environment variables `POKEAPI_URL` and `TRANSLATOR_API_URL` prior to execution.

## How to test

```$ go test ./...```

## Send a request

```$ curl -i localhost:8080/pokemon/translated/mewtwo```
