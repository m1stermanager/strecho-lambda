package main

import (
	"strecho-lambda/internal/handler"
	"strecho-lambda/pkg/echo"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	//todo: routing based on the selected intent etc
	lambda.Start(func(request *echo.Request) (*echo.Response, error) {
		handler := handler.NewGetActivityHandler(request)
		return handler.Handle()
	})
}
