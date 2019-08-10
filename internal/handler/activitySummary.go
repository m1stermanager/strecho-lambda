package handler

import (
	"fmt"
)

type activitySummary interface {
	Summarize(measurementPreference) (string, error)
}

type runningAttributes struct{}
type bikingAttributes struct{}
type swimmingAttrbutes struct{}

func (aggInfo *aggregatedActivityInfo) Summarize(measurementPreference measurementPreference) (string, error) {
	summaryMessage := ""

	activityNoun := aggInfo.activityAttributes.singular
	if aggInfo.activityCount > 1 {
		activityNoun = aggInfo.activityAttributes.plural
	}

	unitOfMeasurement := meters
	if measurementPreference == measurementPreferences.Imperial {
		unitOfMeasurement = aggInfo.activityAttributes.imperialMeasure
	} else {
		unitOfMeasurement = aggInfo.activityAttributes.metricMeasure
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
