package handler

import (
	"fmt"
	"strings"

	"strecho-lambda/internal/client"
	"strecho-lambda/pkg/echo"

	strava "github.com/strava/go.strava"
)

//NewGetActivityHandler takes in a request and gives you a handler
func NewGetActivityHandler(token string) *GetActivityHandler {
	handler := GetActivityHandler{
		client.NewStravaClient(token),
	}

	return &handler
}

//GetActivityHandler takes a request and pulls the last 24 hours of activity in order to
//provide a summary to the user.
type GetActivityHandler struct {
	stravaClient StravaClient
}

//Handle receives an echo request, processes the last 24 hours worth of
//activity for the athlete and provides a semi-personalized response
func (handler *GetActivityHandler) Handle() (*echo.Response, error) {
	activities, err := handler.stravaClient.GetLast24HoursOfActivity()
	if err != nil {
		return echo.NewPlainTextSpeech("There was an issue talking to strava. Try again later."), err
	}

	athlete, err := handler.stravaClient.GetAthlete()
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

	aggregated := aggregateActivities(activities)

	summaryMessage := ""
	for _, activity := range aggregated {
		summaryMessage += fmt.Sprintf("%s, ", activity.Summary(
			measurementPreference(athlete.MeasurementPreference)),
		)
	}
	summaryMessage = strings.Trim(strings.TrimSpace(summaryMessage), ",")

	return fmt.Sprintf("Looks like %s has %s", athlete.FirstName, summaryMessage)
}
