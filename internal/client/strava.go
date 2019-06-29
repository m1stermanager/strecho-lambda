package client

import (
	"time"

	strava "github.com/strava/go.strava"
)

//NewStravaClient returns a pointer to a StravaClient
func NewStravaClient(token string) *StravaClient {
	internalClient := strava.NewClient(token)
	client := StravaClient{
		strava.NewCurrentAthleteService(internalClient),
	}

	return &client
}

//StravaClient is an implementation of handler.StravaClient
type StravaClient struct {
	currentAthleteService *strava.CurrentAthleteService
}

//GetAthlete is an implementation of the thing
func (client *StravaClient) GetAthlete() (*strava.AthleteDetailed, error) {
	return client.currentAthleteService.Get().Do()
}

//GetLast24HoursOfActivity is a thing
func (client *StravaClient) GetLast24HoursOfActivity() ([]*strava.ActivitySummary, error) {
	//blah blah something about 2038
	now := time.Now().UTC()
	beginningOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	endOfDay := beginningOfDay.Add((time.Hour * 24) - time.Second)

	return client.currentAthleteService.
		ListActivities().
		PerPage(1).
		Before(int(endOfDay.Unix())).
		After(int(beginningOfDay.Unix())).
		PerPage(10). //if you have more than 10 activities in a day i'm gonna be impressed
		Page(1).
		Do()
}
