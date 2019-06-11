package echo

//Session is the data that relates to the current "context" of the user's interaction with alexa.
//This is basically the state of the current conversation with alexa
type Session struct {
	New       bool   `json:"new"`
	SessionID string `json:"sessionId"`
}

//Context is the state of the actual alexa device at the time of the request
type Context struct {
	System System `json:"system"`
}

//System is the meat of the context.... the state of the device making the request, the request you are
//receiving, the user the device thinks its coming from and all that jazz.
type System struct {
	AccessToken string `json:"apiAccessToken"`
	Endpoint    string `json:"apiEndpoint"`
}

//Slot is the information about 'slot' being matched
type Slot struct {
	Name               string `json:"name"`
	Value              string `json:"value"`
	ConfirmationStatus string `json:"confirmationStatus"`
}
