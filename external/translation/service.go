package translation

import (
	"encoding/json"
	"log"

	"github.com/lindronics/pokedex/external"
)

const translatorApiUrl string = "https://api.funtranslations.com/translate"

type Translator string

const (
	Shakespeare Translator = "shakespeare"
	Yoda        Translator = "yoda"
)

func TranslateText(text string, translator Translator) string {
	responseBody := external.PostCall(translatorApiUrl, string(translator), text)
	var response Response
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		log.Fatal("Response body corrupted")
	}
	return response.Contents.Translated
}
