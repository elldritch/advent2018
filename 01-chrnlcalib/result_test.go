package chrnlcalib_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	chrnlcalib "github.com/liftM/advent2018/01-chrnlcalib"
)

type ResultTestCase struct {
	Changes        chrnlcalib.Changes
	ExpectedResult int
}

var resultCases = []ResultTestCase{
	ResultTestCase{
		Changes:        chrnlcalib.Changes{1, -2, 3, 1},
		ExpectedResult: 3,
	},
	ResultTestCase{
		Changes:        chrnlcalib.Changes{1, 1, 1},
		ExpectedResult: 3,
	},
	ResultTestCase{
		Changes:        chrnlcalib.Changes{1, 1, -2},
		ExpectedResult: 0,
	},
	ResultTestCase{
		Changes:        chrnlcalib.Changes{-1, -2, -3},
		ExpectedResult: -6,
	},
}

func TestResult(t *testing.T) {
	for _, tc := range resultCases {
		assert.Equal(t, chrnlcalib.Result(tc.Changes), tc.ExpectedResult)
	}
}
