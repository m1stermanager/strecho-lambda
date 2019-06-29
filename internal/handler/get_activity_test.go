package handler

import (
	"strecho-lambda/internal/handler/mock_handler"
	"testing"

	"github.com/golang/mock/gomock"
	strava "github.com/strava/go.strava"
	"github.com/stretchr/testify/assert"
)

func Test_HandlerHappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStrava := mock_handler.NewMockStravaClient(ctrl)

	SUT := new(GetActivityHandler)
	SUT.stravaClient = mockStrava

	activities := []*strava.ActivitySummary{
		fakeActivity(strava.ActivityTypes.Run, 5000),
	}
	athlete := fakeAthlete("test")

	mockStrava.EXPECT().GetAthlete().Return(athlete, nil)
	mockStrava.EXPECT().GetLast24HoursOfActivity().Return(activities, nil)

	response, err := SUT.Handle()
	assert.Nil(t, err)
	assert.NotNil(t, response)

	expectedSpeech := "Looks like test has 1 run for 3.11 miles"
	assert.Equal(t, expectedSpeech, *response.Response.OutputSpeech.Text)
}

func fakeActivity(activityType strava.ActivityType, meters float64) *strava.ActivitySummary {
	activity := new(strava.ActivitySummary)

	activity.Type = activityType
	activity.Distance = meters

	return activity
}

func fakeAthlete(firstName string) *strava.AthleteDetailed {
	athlete := new(strava.AthleteDetailed)
	athlete.FirstName = firstName

	return athlete
}
