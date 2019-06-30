package handler

import (
	"fmt"
	"math"
)

type measurementPreference string

var measurementPreferences = struct {
	Feet   measurementPreference
	Meters measurementPreference
}{"Feet", "Meters"}

type activitySummary interface {
	Summary(measurementPreference) string
}

type runningInfo aggregatedActivityInfo
type bikingInfo aggregatedActivityInfo
type swimmingInfo aggregatedActivityInfo

func (run *runningInfo) Summary(measurementPreference measurementPreference) string {
	summaryMessage := ""

	activityNoun := "run"
	if run.activityCount > 1 {
		activityNoun = "runs"
	}

	if measurementPreference == measurementPreferences.Feet {
		miles := math.Round((run.meters*0.0006213712)*100) / 100
		summaryMessage += fmt.Sprintf("%v %s for %v miles", run.activityCount, activityNoun, miles)
	} else {
		kilometers := math.Round((run.meters/1000)*100) / 100
		summaryMessage += fmt.Sprintf("%v %s for %v kilometers", run.activityCount, activityNoun, kilometers)
	}

	if run.kudos > 0 {
		summaryMessage += fmt.Sprintf("and has received %v kudos", run.kudos)
	}

	return summaryMessage
}

func (bike *bikingInfo) Summary(measurementPreference measurementPreference) string {
	summaryMessage := ""
	activityNoun := "ride"
	if bike.activityCount > 1 {
		activityNoun = "rides"
	}

	if measurementPreference == measurementPreferences.Feet {
		miles := math.Round((bike.meters*0.0006213712)*100) / 100
		summaryMessage += fmt.Sprintf("%v %s for %v miles", bike.activityCount, activityNoun, miles)
	} else {
		kilometers := math.Round((bike.meters*1000)*100) / 100
		summaryMessage += fmt.Sprintf("%v %s for %v kilometers", bike.activityCount, activityNoun, kilometers)
	}

	if bike.kudos > 0 {
		summaryMessage += fmt.Sprintf("and has received %v kudos", bike.kudos)
	}

	return summaryMessage
}

func (swim *swimmingInfo) Summary(measurementPreference) string {
	summaryMessage := ""
	activityNoun := "swim"
	if swim.activityCount > 1 {
		activityNoun = "swims"
	}

	var roundedMeters = math.Round(swim.meters*100) / 100
	summaryMessage += fmt.Sprintf("%v %s for %v kilometers", swim.activityCount, activityNoun, roundedMeters)

	if swim.kudos > 0 {
		summaryMessage += fmt.Sprintf("and has received %v kudos", swim.kudos)
	}

	return summaryMessage
}
