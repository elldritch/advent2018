package main

import (
	"flag"
	"fmt"

	chronalcalibration "github.com/liftM/advent2018/01-chronalcalibration"
)

func main() {
	input := flag.String("input", "", "path to input file")
	flag.Parse()

	if *input == "" {
		panic("-input must be set")
	}

	changes, err := chronalcalibration.ReadFile(*input)
	if err != nil {
		panic(err)
	}

	result := chronalcalibration.Result(changes)
	fmt.Printf("Part one: %d\n", result)

	repeat := chronalcalibration.Repeat(changes)
	fmt.Printf("Part two: %d\n", repeat)
}
