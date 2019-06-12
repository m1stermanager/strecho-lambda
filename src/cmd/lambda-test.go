package main

import (
	"fmt"
	"github/m1stermanager/strecho-lambda/src/echo"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(dummyAlexaHandler)

	fmt.Println("done")
}

func helloWorldHandler() (string, error) {
	return "hello world", nil
}

func dummyAlexaHandler(request *echo.Request) (*echo.Response, error) {
	return echo.NewPlainTextSpeech("hello world"), nil
}
