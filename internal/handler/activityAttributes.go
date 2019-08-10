package handler

type activityAttributes interface {
	singular() string
	plural() string
	metricEquivalent() unitOfMeasurement
	imperialEquivalent() unitOfMeasurement
}

func (*runningAttributes) singular() string {
	return "run"
}

func (*runningAttributes) plural() string {
	return "runs"
}

func (*runningAttributes) metricEquivalent() unitOfMeasurement {
	return kilometers
}

func (*runningAttributes) imperialEquivalent() unitOfMeasurement {
	return miles
}

func (*bikingAttributes) singular() string {
	return "ride"
}

func (*bikingAttributes) plural() string {
	return "rides"
}

func (*bikingAttributes) metricEquivalent() unitOfMeasurement {
	return kilometers
}

func (*bikingAttributes) imperialEquivalent() unitOfMeasurement {
	return miles
}

func (*swimmingAttrbutes) singular() string {
	return "swim"
}

func (*swimmingAttrbutes) plural() string {
	return "swims"
}

func (*swimmingAttrbutes) metricEquivalent() unitOfMeasurement {
	return meters
}

func (*swimmingAttrbutes) imperialEquivalent() unitOfMeasurement {
	return yards
}
