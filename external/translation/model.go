package translation

type Request struct {
	Text string `json:"text"`
}

type Response struct {
	Contents ResponseContents `json:"contents"`
}

type ResponseContents struct {
	Translated  string `json:"translated"`
	Text        string `json:"text"`
	Translation string `json:"translation"`
}
