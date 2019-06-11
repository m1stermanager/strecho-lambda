package echo

//Request is a request specifically geared towards fulfilling a user's request
type Request struct {
	Version string  `json:"version"`
	Session Session `json:"session"`

	Request IntentRequest `json:"request"`
}

//IntentRequest is the body of an IntentRequest request
type IntentRequest struct {
	Type      string `json:"type"`
	RequestID string `json:"requestId"`
	Timestamp string `json:"timestamp"`
	Locale    string `json:"locale"`

	Intent Intent `json:"intent"`
}

//Intent is the specific deatil about the intent
type Intent struct {
	Name               string  `json:"name"`
	ConfirmationStatus string  `json:"confirmationStatus"`
	Slots              []*Slot `json:"slots"`
}
