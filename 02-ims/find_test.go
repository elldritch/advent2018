package ims_test

import "testing"

import ims "github.com/liftM/advent2018/02-ims"

import "github.com/stretchr/testify/assert"

func TestFind(t *testing.T) {
	words := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}

	answers, err := ims.Find(words)
	assert.NoError(t, err)
	assert.Contains(t, answers, "fghij")
	assert.Contains(t, answers, "fguij")
}

func TestCommon(t *testing.T) {
	common := ims.Common("fghij", "fguij")
	assert.Equal(t, common, "fgij")

	common = ims.Common("fonbwmjyquwtapeyzikghtvdxl", "fonbwmjzquwtapeyzikghtvdxl")
	assert.Equal(t, common, "fonbwmjquwtapeyzikghtvdxl")
}
