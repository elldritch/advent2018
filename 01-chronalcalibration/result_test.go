package chronalcalibration_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	chronalcalibration "github.com/liftM/advent2018/01-chronalcalibration"
)

type ResultTestCase struct {
	Changes        chronalcalibration.Changes
	ExpectedResult int
}

var resultCases = []ResultTestCase{
	ResultTestCase{
		Changes:        chronalcalibration.Changes{1, -2, 3, 1},
		ExpectedResult: 3,
	},
	ResultTestCase{
		Changes:        chronalcalibration.Changes{1, 1, 1},
		ExpectedResult: 3,
	},
	ResultTestCase{
		Changes:        chronalcalibration.Changes{1, 1, -2},
		ExpectedResult: 0,
	},
	ResultTestCase{
		Changes:        chronalcalibration.Changes{-1, -2, -3},
		ExpectedResult: -6,
	},
}

func TestResult(t *testing.T) {
	for _, tc := range resultCases {
		assert.Equal(t, chronalcalibration.Result(tc.Changes), tc.ExpectedResult)
	}
}
