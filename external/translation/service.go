package translation

import (
	"github.com/lindronics/pokedex/external"
)

const translatorApiUrl string = "https://api.funtranslations.com/translate"

// Translator reflects different translator options of the funtranslations API
type Translator string

const (
	Shakespeare Translator = "shakespeare"
	Yoda        Translator = "yoda"
)

// TranlateText translates a given text with a given translator setting
// by calling the funtranslations external API.
// If an error occurs, returns nil and an error object containing the status code to return.
func TranslateText(text string, translator Translator) (string, *external.CallError) {
	var response Response
	err := external.PostCall(translatorApiUrl, string(translator), Request{text}, &response)
	return response.Contents.Translated, err
}
