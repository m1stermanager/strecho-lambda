package handler

import (
	"fmt"
	"math"
)

var (
	errUnknownUnitError = fmt.Errorf("unknown unit of measure provided")
)

type unitOfMeasurement int

const (
	meters unitOfMeasurement = iota
	kilometers
	feet
	yards
	miles
)

func (measure unitOfMeasurement) String() string {
	switch measure {
	case meters:
		return "meters"
	case kilometers:
		return "kilometers"
	case feet:
		return "feet"
	case yards:
		return "yards"
	case miles:
		return "miles"
	}

	return "unknown"
}

func convertFromMeters(m float64, convertTo unitOfMeasurement) (float64, error) {
	var converted float64

	switch convertTo {
	case meters:
		converted = m
	case kilometers:
		converted = m / 1000
	case feet:
		converted = m * 3.28084
	case yards:
		converted = (m * 3.28084) / 3
	case miles:
		converted = (m * 3.28084) / 5280
	default:
		return 0, errUnknownUnitError
	}

	rounded := math.Round(converted*100) / 100

	return rounded, nil
}
