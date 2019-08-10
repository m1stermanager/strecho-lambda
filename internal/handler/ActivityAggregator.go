package handler

import strava "github.com/strava/go.strava"

type aggregatedActivityInfo struct {
	activityCount int
	meters        float64
	kudos         int
	commentCount  int

	activityAttributes activityAttrs
}

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
