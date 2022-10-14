package isbn

import (
	"regexp"
	"strconv"
)

func IsValidISBN(isbn string) bool {
	if len(regexp.MustCompile(`[^\dX-]`).FindAllString(isbn, -1)) > 0 {
		return false
	}
	nlist := regexp.MustCompile(`[\dX]`).FindAllString(isbn, -1)
	count := len(nlist)
	mod := count + 1
	if count != 10 && count != 13 {
		return false
	}
	t := 0
	mult := len(nlist)
	for _, i := range nlist {
		switch {
		case mult > 1:
			digit, err := strconv.Atoi(i)
			if err != nil {
				return false
			}
			value := digit * mult
			t += value
			mult--
		case mult == 1 && i != "X":
			digit, _ := strconv.Atoi(i)
			t += digit
		case mult == 1 && i == "X":
			t += 10
		}
	}
	return t%mod == 0
}
