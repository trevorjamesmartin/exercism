package atbash

import "strings"

func Atbash(s string) string {
	var result string
	var charCount int
	max := len(s) - 1
	for i, r := range s {
		switch {
		case r > 64 && r < 91:
			c := 25 - (r - 65)
			result += string(rune(c + 65 + 32))
			charCount++
		case r > 96 && r < 123:
			c := 25 - (r - 97)
			result += string(rune(c + 97))
			charCount++
		case r > 47 && r < 58:
			result += string(rune(r))
			charCount++
		}
		if charCount == 5 && i < max {
			result += " "
			charCount = 0
		}
	}
	return strings.TrimSpace(result)
}
