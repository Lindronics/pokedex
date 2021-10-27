// Contains client service and models for the funtranslations external API
package translation

import (
	"os"

	"github.com/lindronics/pokedex/external"
)

const TranslatorApiUrlParam string = "TRANSLATOR_API_URL"

// Provider translates a given string using a given TranslatorType
type Provider interface {
	TranslateText(string, TranslatorType) (string, *external.CallError)
}

type HttpProvider struct {
	BaseUrl string
}

func NewHttpProvider() *HttpProvider {
	return &HttpProvider{os.Getenv(TranslatorApiUrlParam)}
}

// TranslateText translates a given text with a given translator setting
// by calling the funtranslations external API.
// If an error occurs, returns nil and an error object containing the status code to return.
func (p *HttpProvider) TranslateText(text string, translator TranslatorType) (string, *external.CallError) {
	var response Response
	err := external.PostCall(p.BaseUrl, string(translator), Request{text}, &response)
	return response.Contents.Translated, err
}
