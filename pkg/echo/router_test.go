package echo

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var routerTestCases = []struct {
	availableKeys  []string
	executedKey    string
	executionError error
}{
	{[]string{"1", "2"}, "2", nil},
	{[]string{"1", "2"}, "3", &ErrUnknownRoute{}},
}

func Test_Execute(t *testing.T) {
	for i, test := range routerTestCases {
		testMessage := fmt.Sprintf("test case #%v", i+1)

		//create a router where each handler returns a response, track the expected response
		//in order to assert that the instance received is what was expected...
		routerMap := make(map[string]Handler)
		responses := make(map[string]*Response)
		for _, key := range test.availableKeys {
			response := new(Response)
			routerMap[key] = func(*Request) (*Response, error) {
				return response, nil
			}
			responses[key] = response
		}
		testRequest := Request{Request: IntentRequest{Intent: Intent{Name: test.executedKey}}}
		router := NewRouter(routerMap)
		response, executionError := router.Execute(&testRequest)
		if test.executionError == nil {
			assert.Nil(t, executionError)
			assert.Same(t, responses[test.executedKey], response)
		} else {
			assert.Nil(t, response)
			assert.True(t, assert.IsType(t, test.executionError, executionError), testMessage)
		}
	}
}

var unknownRouteTestCases = []struct {
	keys          []string
	attempted     string
	expectedError string
}{
	{[]string{"a", "b", "c"}, "d", "attempted to use key 'd' but only available keys are 'a, b, c'"},
	{[]string{}, "a", "attempted to use key 'a' but only available keys are ''"},
	{[]string{"a"}, "b", "attempted to use key 'b' but only available keys are 'a'"},
}

func Test_ErrUnknownRoute_String(t *testing.T) {
	for i, test := range unknownRouteTestCases {
		testMessage := fmt.Sprintf("test case #%v", i+1)

		err := ErrUnknownRoute{AvailableKeys: test.keys, AttemptedKey: test.attempted}
		errText := err.Error()
		assert.Equal(t, test.expectedError, errText, testMessage)
	}
}

var isErrorUnknownRouteTestCases = []struct {
	err      error
	response bool
}{
	{new(ErrUnknownRoute), true},
	{errors.New("test"), false},
	{nil, false},
}

func Test_IsErrorUnknownRoute(t *testing.T) {
	for i, test := range isErrorUnknownRouteTestCases {
		testMessage := fmt.Sprintf("test case #%v", i+1)

		response := IsErrorUnknownRoute((test.err))
		assert.Equal(t, test.response, response, testMessage)
	}
}
