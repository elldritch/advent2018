package chrnlcalib

import (
	"errors"
	"strconv"
)

type Changes []int

func Parse(lines []string) (Changes, error) {
	var changes Changes
	for _, line := range lines {
		sign := 1

		if line[0] == '+' {
			sign = 1
		} else if line[0] == '-' {
			sign = -1
		} else {
			return nil, errors.New("malformed line")
		}

		num, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}
		changes = append(changes, num*sign)
	}

	return changes, nil
}

func Result(changes Changes) int {
	frequency := 0
	for _, change := range changes {
		frequency += change
	}
	return frequency
}

func Repeat(changes Changes) int {
	seen := make(map[int]bool)
	frequency := 0
	for {
		for _, change := range changes {
			if _, ok := seen[frequency]; ok {
				return frequency
			}
			seen[frequency] = true

			frequency += change
		}
	}
}
