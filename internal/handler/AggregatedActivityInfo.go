package handler

import (
	"fmt"

	strava "github.com/strava/go.strava"
)

type aggregatedActivityInfo struct {
	activityCount int
	meters        float64
	kudos         int
	commentCount  int

	activityAttributes activityAttributes
}

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

	//i.e. 2 runs for 20 kilometers
	summaryMessage += fmt.Sprintf("%v %s for %v %s", aggInfo.activityCount, activityNoun, convertedUnits, unitOfMeasurement)

	if aggInfo.kudos > 0 {
		summaryMessage += fmt.Sprintf("and has received %v kudos", aggInfo.kudos)
	}

	return summaryMessage, nil
}

//summarizeActivities receives activities from strava and then summarizes them by aggregated various
//points of data together. The result is a collection of summarized activities. One result per
//type of activity will be returned. If an unrecognized or unsupported activity type is given,
//it will be skipped but will not cause the call to fail
func summarizeActivities(activities []*strava.ActivitySummary) []activitySummary {
	activityTypeMap := make(map[strava.ActivityType]*aggregatedActivityInfo)
	summaries := make([]activitySummary, 0)

	for _, activity := range activities {
		attrs, exists := activityAttrMap[activity.Type]
		if !exists {
			//we don't support this activity type.... but we also don't want to explode
			continue
		}

		agg, exists := activityTypeMap[activity.Type]
		if !exists {
			freshAggregation := aggregatedActivityInfo{activityAttributes: *attrs}
			agg = &freshAggregation
			activityTypeMap[activity.Type] = agg
			summaries = append(summaries, agg)
		}

		agg.activityCount++
		agg.kudos += activity.KudosCount
		agg.meters += activity.Distance
		agg.commentCount += activity.CommentCount
	}

	return summaries
}
