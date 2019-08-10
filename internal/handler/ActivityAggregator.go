package handler

import strava "github.com/strava/go.strava"

type aggregatedActivityInfo struct {
	activityCount int
	meters        float64
	kudos         int
	commentCount  int

	activityAttributes activityAttributes
}

func summarizeActivities(activities []*strava.ActivitySummary) []activitySummary {
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
			activity.activityAttributes = &runningAttributes{}
		case strava.ActivityTypes.Ride:
			activity.activityAttributes = &bikingAttributes{}
		case strava.ActivityTypes.Swim:
			activity.activityAttributes = &swimmingAttrbutes{}
		}

		summaries = append(summaries, activity)
	}

	return summaries
}
