package main

import (
	"strecho-lambda/internal/handler"
	"strecho-lambda/pkg/echo"

	"github.com/aws/aws-lambda-go/lambda"
)

var router echo.Router

func init() {
	handlers := map[string]echo.Handler{
		"Activity_Today_Summary": handler.NewGetActivityHandler(),
	}

	router = echo.NewRouter(handlers)
}

func main() {
	lambda.Start(func(request *echo.Request) (*echo.Response, error) {
		return router.Execute(request)
	})
}
