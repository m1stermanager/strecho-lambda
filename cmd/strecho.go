package main

import (
	"fmt"
	"strecho-lambda/internal/handler"
	"strecho-lambda/pkg/echo"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	//todo: routing based on the selected intent etc
	lambda.Start(func(request *echo.Request) (*echo.Response, error) {
		token := request.Context.System.User.AccessToken
		handler := handler.NewGetActivityHandler(token)
		return handler.Handle()
	})

	fmt.Println("done")
}
