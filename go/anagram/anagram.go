package anagram

import (
	"reflect"
	"strings"
)

func CountRunes(s string) map[rune]int {
	count := map[rune]int{}
	for _, r := range strings.ToLower(s) {
		count[r]++
	}
	return count
}

func Detect(subject string, candidates []string) []string {
	var result []string
	a := CountRunes(subject)
	for _, c := range candidates {
		b := CountRunes(c)
		if reflect.DeepEqual(a, b) && !strings.EqualFold(c, subject) {
			result = append(result, c)
		}
	}
	return result
}
