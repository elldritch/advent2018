package main

import (
	"flag"
	"fmt"

	chrnlcalib "github.com/liftM/advent2018/01-chrnlcalib"
	"github.com/liftM/advent2018/inpututil"
)

func main() {
	input := flag.String("input", "", "path to input file")
	flag.Parse()

	if *input == "" {
		panic("-input must be set")
	}

	lines, err := inpututil.ReadFile(*input)
	if err != nil {
		panic(err)
	}
	changes, err := chrnlcalib.Parse(lines)
	if err != nil {
		panic(err)
	}

	result := chrnlcalib.Result(changes)
	fmt.Printf("Part one: %d\n", result)

	repeat := chrnlcalib.Repeat(changes)
	fmt.Printf("Part two: %d\n", repeat)
}
