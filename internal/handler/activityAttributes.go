package handler

import strava "github.com/strava/go.strava"

type activityAttrs struct {
	singular        string
	plural          string
	metricMeasure   unitOfMeasurement
	imperialMeasure unitOfMeasurement
}

var activityAttrMap = map[strava.ActivityType]*activityAttrs{
	strava.ActivityTypes.Run: &activityAttrs{
		singular:        "run",
		plural:          "runs",
		metricMeasure:   kilometers,
		imperialMeasure: miles,
	},
	strava.ActivityTypes.Ride: &activityAttrs{
		singular:        "ride",
		plural:          "rides",
		metricMeasure:   kilometers,
		imperialMeasure: miles,
	},
	strava.ActivityTypes.Swim: &activityAttrs{
		singular:        "swim",
		plural:          "swims",
		metricMeasure:   meters,
		imperialMeasure: yards,
	},
}
