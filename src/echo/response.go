package echo

//Response represents an amazon alexa/echo response.
//created per https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html
type Response struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech"`
	ShouldEndSession bool         `json:"shouldEndSession"`
}

//OutputSpeech is the details of the speech
type OutputSpeech struct {
	Type string  `json:"type"`
	Text *string `json:"text"`
	Ssml *string `json:"ssml"`
}

//NewPlainTextSpeech takes words and produces a valid OutputSpeech response. It will end the session.
func NewPlainTextSpeech(words string) *Response {
	speech := OutputSpeech{
		Type: "PlainText",
		Text: &words,
		Ssml: nil,
	}

	response := Response{
		OutputSpeech:     speech,
		ShouldEndSession: true,
	}

	return &response
}
