package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	result := map[string]int{}
	for value, keys := range in {
		for _, i := range keys {
			result[strings.ToLower(i)] = value
		}
	}
	return result
}
