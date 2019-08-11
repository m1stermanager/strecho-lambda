package handler

import strava "github.com/strava/go.strava"

//StravaClient defines the operations we expect to be performing against strava....
//the goal is to essentially organize the operations that the handlers require into
//logical units and abstract the provided strava api away at least somewhat
type StravaClient interface {
	GetAthlete() (*strava.AthleteDetailed, error)
	GetLast24HoursOfActivity() ([]*strava.ActivitySummary, error)
}

type measurementPreference string

//strava shows this preference as feet/meters....
//here we're just translating it to make *maybe*
//more sense generically
var measurementPreferences = struct {
	Imperial measurementPreference
	Metric   measurementPreference
}{"feet", "meters"}
