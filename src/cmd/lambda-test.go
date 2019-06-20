package main

import (
	"fmt"
	"github/m1stermanager/strecho-lambda/src/echo"

	"github.com/aws/aws-lambda-go/lambda"
	strava "github.com/strava/go.strava"
)

func main() {
	lambda.Start(dummyAlexaHandler)

	fmt.Println("done")
}

func helloWorldHandler() (string, error) {
	return "hello world", nil
}

func dummyAlexaHandler(request *echo.Request) (*echo.Response, error) {
	token := request.Context.System.User.AccessToken
	fmt.Println("access token", token)

	client := strava.NewClient(token)
	athService := strava.NewCurrentAthleteService(client)

	athlete := athService.Get().Do()
	if err != nil {
		fmt.Println("error encountered:", err)
		return echo.NewPlainTextSpeech("whoooooops"), nil
	}

	return echo.NewPlainTextSpeech("Hello " + athlete.FirstName), nil
}
