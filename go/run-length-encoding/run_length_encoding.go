package encode

import (
	"fmt"
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	if len(input) == 0 {
		return input
	}
	var output string
	list := []rune(input)
	holding := list[0]
	count := 1
	for index := 1; index < len(list); index++ {
		current := list[index]
		switch {
		case holding == current:
			count++
		case holding != current && count > 1:
			output += fmt.Sprintf("%v%v", count, string(holding))
			holding = current
			count = 1
		case holding != current && count == 1:
			output += string(holding)
			holding = current
			count = 1
		}
	}
	if count > 1 {
		output += fmt.Sprintf("%v", count)
	}
	return output + string(holding)
}

func RunLengthDecode(input string) string {
	var output string
	var multr []rune
	for _, r := range input {
		s := string(multr)
		switch {
		case r > 47 && r < 58:
			multr = append(multr, r)
		default:
			m, _ := strconv.Atoi(s)
			if m > 0 {
				output += strings.Repeat(string(r), m)
				multr = nil
			} else {
				output += string(r)
			}
		}
	}
	return output
}
