package translation

type Request struct {
	Text string `json:"text"`
}

type Response struct {
	Contents ResponseContents `json:"contents" valid:"required"`
}

type ResponseContents struct {
	Translated  string `json:"translated" valid:"required"`
	Text        string `json:"text" valid:"required"`
	Translation string `json:"translation" valid:"required"`
}
