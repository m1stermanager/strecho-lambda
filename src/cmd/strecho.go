package main

import (
	"fmt"
	"github/m1stermanager/strecho-lambda/src/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	//todo: routing based on the selected intent etc
	lambda.Start(handler.GetActivityHandler)

	fmt.Println("done")
}
