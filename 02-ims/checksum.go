package ims

type WordCount struct {
	HasTwo   bool
	HasThree bool
}

func Count(word string) WordCount {
	counts := make(map[rune]int)
	for _, char := range word {
		counts[char]++
	}

	var wc WordCount
	for _, count := range counts {
		if count == 2 {
			wc.HasTwo = true
		}
		if count == 3 {
			wc.HasThree = true
		}
	}

	return wc
}

func Counts(words []string) []WordCount {
	var counts []WordCount
	for _, w := range words {
		counts = append(counts, Count(w))
	}

	return counts
}

func Checksum(counts []WordCount) int {
	twos := 0
	threes := 0

	for _, c := range counts {
		if c.HasTwo {
			twos++
		}
		if c.HasThree {
			threes++
		}
	}

	return twos * threes
}
