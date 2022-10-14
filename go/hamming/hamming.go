package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance calculates the Hamming distance between two DNA strands
func Distance(a, b string) (int, error) {
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return 0, errors.New("sequences are different lengths")
	}
	diff := 0
	for i := range a {
		if b[i] != a[i] {
			diff++
		}
	}
	return diff, nil
}
