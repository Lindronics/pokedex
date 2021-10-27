// Contains client service and models for the funtranslations external API
package translation

import (
	"os"

	"github.com/lindronics/pokedex/external"
)

const translatorApiUrlParam string = "TRANSLATOR_API_URL"

// TranlateText translates a given text with a given translator setting
// by calling the funtranslations external API.
// If an error occurs, returns nil and an error object containing the status code to return.
func TranslateText(text string, translator Translator) (string, *external.CallError) {
	var response Response
	err := external.PostCall(os.Getenv(translatorApiUrlParam), string(translator), Request{text}, &response)
	return response.Contents.Translated, err
}
