package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	vacuum := strings.ReplaceAll(id, " ", "") // Spaces are allowed in the input, but they should be stripped before checking
	pattern := `[0-9]`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(id, -1)
	var numeric = ""
	for _, digit := range matches {
		// (starting from the right)
		numeric = digit + numeric
	}

	if len(vacuum) < 2 || len(vacuum) > len(numeric) {
		// Strings of length 1 or less are not valid
		// All other non-digit characters are disallowed
		return false
	}

	var sum int64

	for i, c := range strings.Split(numeric, "") {
		n, err := strconv.ParseInt(c, 10, 0)
		if err != nil {
			return false
		}
		var digit = n
		if (i+1)%2 == 0 {
			// double every second digit
			dbl := digit * 2
			if dbl > 9 {
				// If doubling the number results in a number greater than 9 then subtract 9 from the product.
				dbl -= 9
			}
			digit = dbl
		}
		sum += digit // sum all of the digits
	}
	// If the sum is evenly divisible by 10, then the number is valid. This number is valid!
	return sum%10 == 0
}
