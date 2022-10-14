package raindrops

import "strconv"

// Convert, given a number, will return a string containing raindrop sounds.
func Convert(number int) string {
	var s string
	if number%3 == 0 {
		s += "Pling"
	}
	if number%5 == 0 {
		s += "Plang"
	}
	if number%7 == 0 {
		s += "Plong"
	}
	if s == "" {
		s = strconv.Itoa(number)
	}
	return s
}
