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

```$ curl -i localhost:8080/pokemon/mewtwo```

```$ curl -i localhost:8080/pokemon/translated/mewtwo```

## Comments

* I omitted some unit tests that would have had a similar structure to the ones I implemented. This is simply due to time constraints. In a production environment, I would probably create a proper testing package containing helpers for easier test construction.
* Go is not my primary language, but I decided that this would be a good opportunity to practice it. I chose it because of its type safety, ecosystem and performance. The code can be somewhat verbose, but this is to increase readability and is idiomatic for this particular language.
* I decided not to set up a full CD pipeline, as it is probably beyond the scope of this exercise. I did create a job to build and run tests for all pushed commits.
