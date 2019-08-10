package handler

import (
	"fmt"
)

type measurementPreference string

//strava shows this preference as feet/meters....
//here we're just translating it to make *maybe*
//more sense generically
var measurementPreferences = struct {
	Imperial measurementPreference
	Metric   measurementPreference
}{"Feet", "Meters"}

type activitySummary interface {
	singular() string
	plural() string
	metricEquivalent() unitOfMeasurement
	imperialEquivalent() unitOfMeasurement

	Summarize(measurementPreference) (string, error)
}

type runningInfo struct{ aggregatedActivityInfo }
type bikingInfo struct{ aggregatedActivityInfo }
type swimmingInfo struct{ aggregatedActivityInfo }

func (*runningInfo) singular() string {
	return "run"
}

func (*runningInfo) plural() string {
	return "runs"
}

func (*runningInfo) metricEquivalent() unitOfMeasurement {
	return kilometers
}

func (*runningInfo) imperialEquivalent() unitOfMeasurement {
	return miles
}

func (*bikingInfo) singular() string {
	return "ride"
}

func (*bikingInfo) plural() string {
	return "rides"
}

func (*bikingInfo) metricEquivalent() unitOfMeasurement {
	return kilometers
}

func (*bikingInfo) imperialEquivalent() unitOfMeasurement {
	return miles
}

func (*swimmingInfo) singular() string {
	return "swim"
}

func (*swimmingInfo) plural() string {
	return "swims"
}

func (*swimmingInfo) metricEquivalent() unitOfMeasurement {
	return meters
}

func (*swimmingInfo) imperialEquivalent() unitOfMeasurement {
	return yards
}

func (aggInfo *aggregatedActivityInfo) Summarize(measurementPreference measurementPreference) (string, error) {
	summaryMessage := ""

	activityNoun := aggInfo.summaryType.singular()
	if aggInfo.activityCount > 1 {
		activityNoun = aggInfo.summaryType.plural()
	}

	unitOfMeasurement := meters
	if measurementPreference == measurementPreferences.Imperial {
		unitOfMeasurement = aggInfo.summaryType.imperialEquivalent()
	} else {
		unitOfMeasurement = aggInfo.summaryType.metricEquivalent()
	}

	convertedUnits, err := convertFromMeters(aggInfo.meters, unitOfMeasurement)
	if err != nil {
		return "", err
	}

	summaryMessage += fmt.Sprintf("%v %s for %v %s", aggInfo.activityCount, activityNoun, convertedUnits, unitOfMeasurement)

	if aggInfo.kudos > 0 {
		summaryMessage += fmt.Sprintf("and has received %v kudos", aggInfo.kudos)
	}

	return summaryMessage, nil
}
