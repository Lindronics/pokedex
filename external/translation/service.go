package translation

import (
	"github.com/lindronics/pokedex/external"
)

const translatorApiUrl string = "https://api.funtranslations.com/translate"

type Translator string

const (
	Shakespeare Translator = "shakespeare"
	Yoda        Translator = "yoda"
)

func TranslateText(text string, translator Translator) (string, *external.HttpError) {
	var response Response
	err := external.PostCall(translatorApiUrl, string(translator), Request{text}, &response)
	return response.Contents.Translated, err
}
