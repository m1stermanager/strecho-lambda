package handler

import strava "github.com/strava/go.strava"

type activityAttributes struct {
	singular        string
	plural          string
	metricMeasure   unitOfMeasurement
	imperialMeasure unitOfMeasurement
}

var activityAttrMap = map[strava.ActivityType]*activityAttributes{
	strava.ActivityTypes.Run: &activityAttributes{
		singular:        "run",
		plural:          "runs",
		metricMeasure:   kilometers,
		imperialMeasure: miles,
	},
	strava.ActivityTypes.Ride: &activityAttributes{
		singular:        "ride",
		plural:          "rides",
		metricMeasure:   kilometers,
		imperialMeasure: miles,
	},
	strava.ActivityTypes.Swim: &activityAttributes{
		singular:        "swim",
		plural:          "swims",
		metricMeasure:   meters,
		imperialMeasure: yards,
	},
}
