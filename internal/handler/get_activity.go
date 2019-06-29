package handler

import (
	"fmt"
	"math"

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

type aggregateActivityInfo struct {
	activityCount int
	meters        float64
	kudos         int
}

func generateActivityStatement(athlete *strava.AthleteDetailed, activities []*strava.ActivitySummary) string {
	if len(activities) == 0 {
		return "Hmm. I'm not seeing any activity for today"
	}

	//an accumulator
	activityTypeDistance := make(map[strava.ActivityType]*aggregateActivityInfo)
	for _, activity := range activities {
		agg, exists := activityTypeDistance[activity.Type]
		if !exists {
			agg = new(aggregateActivityInfo)
			activityTypeDistance[activity.Type] = agg
		}
		agg.activityCount++
		agg.kudos += activity.KudosCount
		agg.meters += activity.Distance
	}

	summaryMessage := ""
	for activityType, distance := range activityTypeDistance {
		pastTenseType := activityType.String() //default to activity type string
		if activityType == strava.ActivityTypes.Run {
			if distance.activityCount > 1 {
				pastTenseType = "runs"
			} else {
				pastTenseType = "run"
			}
		}

		miles := math.Round((distance.meters*0.0006213712)*100) / 100
		summaryMessage += fmt.Sprintf("has %v %s for %v miles", distance.activityCount, pastTenseType, miles)
		if distance.kudos > 0 {
			summaryMessage += fmt.Sprintf(" and has received %v kudos", distance.kudos)
		}
	}

	return fmt.Sprintf("Looks like %s %s", athlete.FirstName, summaryMessage)
}
