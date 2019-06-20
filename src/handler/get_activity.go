package handler

import (
	"fmt"
	"github/m1stermanager/strecho-lambda/src/client"
	"github/m1stermanager/strecho-lambda/src/echo"

	strava "github.com/strava/go.strava"
)

//GetActivityHandler receives an echo request, processes the last 24 hours worth of
//activity for the athlete and provides a semi-personalized response
func GetActivityHandler(request *echo.Request) (*echo.Response, error) {
	token := request.Context.System.User.AccessToken
	//todo obviously work this out into the main package and allow the handler to 
	//receive the interface and not instantiate the type directly
	stravaClient := client.NewStravaClient(token)

	activities, err := stravaClient.GetLast24HoursOfActivity()
	if err != nil {
		return echo.NewPlainTextSpeech("There was an issue talking to strava. Try again later."), err
	}

	athlete, err := stravaClient.GetAthlete()
	if err != nil {
		return echo.NewPlainTextSpeech("There was an issue talking to strava. Try again later."), err
	}

	speech := generateActivityStatement(athlete, activities)
	return echo.NewPlainTextSpeech(speech), nil
}

func generateActivityStatement(athlete *strava.AthleteDetailed, activities []*strava.ActivitySummary) string {
	if len(activities) == 0 {
		return "Hmm. I'm not seeing any activity for today"
	}

	//an accumulator
	activityTypeDistance := make(map[strava.ActivityType]float64)
	for _, activity := range activities {
		activityTypeDistance[activity.Type] += activity.Distance
	}

	return fmt.Sprintln(athlete.FirstName, ", I'm seeing", len(activities), "activities in the last 24 hours")
}
