package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var meterConversionTests = []struct {
	meters         float64
	targetUnits    unitOfMeasurement
	expectedResult float64
	expectedError  error
}{
	{5000, miles, 3.11, nil},
	{5000, meters, 5000, nil},
	{5000, kilometers, 5, nil},
	{5000, feet, 16404.2, nil},
	{5000, yards, 5468.07, nil},
	{5000, unitOfMeasurement(-100), 0, unknownUnitError},
}

func Test_ConvertFromMeters(t *testing.T) {
	for i, test := range meterConversionTests {
		converted, err := convertFromMeters(test.meters, test.targetUnits)
		assert.Equal(t, test.expectedResult, converted, "test case %v", i+1)
		assert.Equal(t, test.expectedError, err, "test case %v", i+1)
	}
}
