package echo

import (
	"fmt"
	"strings"
)

//ErrUnknownRoute is used when the intent name in the request is not registered with the router
type ErrUnknownRoute struct {
	AvailableKeys []string
	AttemptedKey  string
}

func (err *ErrUnknownRoute) Error() string {
	keys := strings.Join(err.AvailableKeys, ", ")
	return fmt.Sprintf("attempted to use key '%v' but only available keys are '%v'", err.AttemptedKey, keys)
}

//IsErrorUnknownRoute will tell you if an error is an instance of ErrUnkownRoute
func IsErrorUnknownRoute(err error) bool {
	_, ok := err.(*ErrUnknownRoute)
	return ok
}

//Handler is the function signature that the router knows how to handle
type Handler func(*Request) (*Response, error)

//Router will drive the selection of an appropriate handler based on the name of an intent
type Router interface {
	Execute(*Request) (*Response, error)
}

//NewRouter takes in an intent/handler map and returns a properly setup Router
func NewRouter(routes map[string]Handler) Router {
	return router(routes)
}

type router map[string]Handler

func (router router) Execute(request *Request) (*Response, error) {
	intentName := request.Request.Intent.Name
	handler, ok := router[intentName]
	if !ok {
		keys := router.getKeys()
		err := ErrUnknownRoute{AvailableKeys: keys, AttemptedKey: intentName}
		return nil, &err
	}
	return handler(request)
}

func (router router) getKeys() []string {
	keys := make([]string, len(router))
	i := 0
	for k := range router {
		keys[i] = k
		i++
	}

	return keys
}
