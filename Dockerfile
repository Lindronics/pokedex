FROM golang:1.17

WORKDIR /go/src/pokedex
COPY . .

ENV POKEAPI_URL="https://pokeapi.co/api/v2"
ENV TRANSLATOR_API_URL="https://api.funtranslations.com/translate"
ENV GIN_MODE="release"

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["pokedex"]
