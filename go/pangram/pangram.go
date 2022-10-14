package pangram

import "strings"

func IsPangram(input string) bool {
	abc := map[rune]int{}
	total := 0
	for _, i := range strings.ToLower(input) {
		if i > 96 && i < 123 && abc[i] == 0 {
			abc[i] = 1
			total++
		}
	}
	return total == 26
}
