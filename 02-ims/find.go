package ims

import "errors"

func Find(words []string) ([2]string, error) {
	l := len(words[0])
	for i, w := range words {
		for j, w2 := range words {
			if i >= j {
				continue
			}

			same := 0
			for c := range w {
				if w[c] == w2[c] {
					same++
				}
			}
			if same == l-1 {
				return [2]string{w, w2}, nil
			}
		}
	}
	return [2]string{}, errors.New("pair not found")
}

func Common(a, b string) string {
	seen := make(map[int]rune)
	for i, c := range a {
		seen[i] = c
	}

	var common string
	for i, c := range b {
		if seen[i] == c {
			common += string(c)
		}
	}

	return common
}
