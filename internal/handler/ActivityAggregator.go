package handler

import strava "github.com/strava/go.strava"

type aggregatedActivityInfo struct {
	activityCount int
	meters        float64
	kudos         int
	commentCount  int
}

func aggregateActivities(activities []*strava.ActivitySummary) []activitySummary {
	activityTypeMap := make(map[strava.ActivityType]*aggregatedActivityInfo)

	for _, activity := range activities {
		agg, exists := activityTypeMap[activity.Type]
		if !exists {
			agg = new(aggregatedActivityInfo)
			activityTypeMap[activity.Type] = agg
		}

		agg.activityCount++
		agg.kudos += activity.KudosCount
		agg.meters += activity.Distance
		agg.commentCount += activity.CommentCount
	}

	summaries := make([]activitySummary, 0)
	for activityType, activity := range activityTypeMap {
		switch activityType {
		case strava.ActivityTypes.Run:
			run := runningInfo(*activity)
			summaries = append(summaries, &run)
		case strava.ActivityTypes.Ride:
			ride := bikingInfo(*activity)
			summaries = append(summaries, &ride)
		case strava.ActivityTypes.Swim:
			swim := swimmingInfo(*activity)
			summaries = append(summaries, &swim)
		}
	}

	return summaries
}
