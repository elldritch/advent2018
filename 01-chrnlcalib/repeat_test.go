package chrnlcalib_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	chrnlcalib "github.com/liftM/advent2018/01-chrnlcalib"
)

type RepeatTestCase struct {
	Changes        chrnlcalib.Changes
	ExpectedRepeat int
}

var cases = []RepeatTestCase{
	RepeatTestCase{
		Changes:        chrnlcalib.Changes{1, -2, 3, 1},
		ExpectedRepeat: 2,
	},
	RepeatTestCase{
		Changes:        chrnlcalib.Changes{1, -1},
		ExpectedRepeat: 0,
	},
	RepeatTestCase{
		Changes:        chrnlcalib.Changes{3, 3, 4, -2, -4},
		ExpectedRepeat: 10,
	},
	RepeatTestCase{
		Changes:        chrnlcalib.Changes{-6, 3, 8, 5, -6},
		ExpectedRepeat: 5,
	},
	RepeatTestCase{
		Changes:        chrnlcalib.Changes{7, 7, -2, -7, -4},
		ExpectedRepeat: 14,
	},
}

func TestRepeat(t *testing.T) {
	for _, tc := range cases {
		assert.Equal(t, chrnlcalib.Repeat(tc.Changes), tc.ExpectedRepeat)
	}
}
