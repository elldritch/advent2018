package main

import (
	"flag"
	"fmt"

	chrnlcalib "github.com/liftM/advent2018/01-chrnlcalib"
)

func main() {
	input := flag.String("input", "", "path to input file")
	flag.Parse()

	if *input == "" {
		panic("-input must be set")
	}

	changes, err := chrnlcalib.ReadFile(*input)
	if err != nil {
		panic(err)
	}

	result := chrnlcalib.Result(changes)
	fmt.Printf("Part one: %d\n", result)

	repeat := chrnlcalib.Repeat(changes)
	fmt.Printf("Part two: %d\n", repeat)
}
