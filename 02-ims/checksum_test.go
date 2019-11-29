package ims_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ims "github.com/liftM/advent2018/02-ims"
)

func TestCounts(t *testing.T) {
	words := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}

	counts := []ims.WordCount{
		ims.WordCount{},
		ims.WordCount{
			HasTwo:   true,
			HasThree: true,
		},
		ims.WordCount{
			HasTwo: true,
		},
		ims.WordCount{
			HasThree: true,
		},
		ims.WordCount{
			HasTwo: true,
		},
		ims.WordCount{
			HasTwo: true,
		},
		ims.WordCount{
			HasThree: true,
		},
	}

	assert.Equal(t, len(words), len(counts))
	for i, w := range words {
		assert.Equal(t, counts[i], ims.Count(w))
	}
}
