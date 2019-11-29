package main

import (
	"flag"
	"fmt"

	ims "github.com/liftM/advent2018/02-ims"
)

func main() {
	input := flag.String("input", "", "path to input file")
	flag.Parse()

	if *input == "" {
		panic("-input must be set")
	}

	lines, err := ims.ReadFile(*input)
	if err != nil {
		panic(err)
	}

	counts := ims.Counts(lines)
	sum := ims.Checksum(counts)

	fmt.Printf("Part one: %d\n", sum)
}
