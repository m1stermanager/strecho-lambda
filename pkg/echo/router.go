package echo

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
	handler := router[request.Request.Intent.Name]
	return handler(request)
}
