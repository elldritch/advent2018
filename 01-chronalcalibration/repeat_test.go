package chronalcalibration_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	chronalcalibration "github.com/liftM/advent2018/01-chronalcalibration"
)

type RepeatTestCase struct {
	Changes        chronalcalibration.Changes
	ExpectedRepeat int
}

var cases = []RepeatTestCase{
	RepeatTestCase{
		Changes:        chronalcalibration.Changes{1, -2, 3, 1},
		ExpectedRepeat: 2,
	},
	RepeatTestCase{
		Changes:        chronalcalibration.Changes{1, -1},
		ExpectedRepeat: 0,
	},
	RepeatTestCase{
		Changes:        chronalcalibration.Changes{3, 3, 4, -2, -4},
		ExpectedRepeat: 10,
	},
	RepeatTestCase{
		Changes:        chronalcalibration.Changes{-6, 3, 8, 5, -6},
		ExpectedRepeat: 5,
	},
	RepeatTestCase{
		Changes:        chronalcalibration.Changes{7, 7, -2, -7, -4},
		ExpectedRepeat: 14,
	},
}

func TestRepeat(t *testing.T) {
	for _, tc := range cases {
		assert.Equal(t, chronalcalibration.Repeat(tc.Changes), tc.ExpectedRepeat)
	}
}
