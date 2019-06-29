package echo

//Response represents an amazon alexa/echo response.
//created per https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html
type Response struct {
	Version  string         `json:"version"`
	Response SpeechResponse `json:"response"`
}

//SpeechResponse is the specific part of the response that makes the echo talk
type SpeechResponse struct {
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
		Version: "1.00",
		Response: SpeechResponse{
			OutputSpeech:     speech,
			ShouldEndSession: true,
		},
	}

	return &response
}
