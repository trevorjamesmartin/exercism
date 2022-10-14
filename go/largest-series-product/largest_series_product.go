package lsproduct

import (
	"errors"
	"regexp"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	var product int64
	switch {
	case span > len(digits):
		return 1, errors.New("span greater than input")
	case span < 0:
		return 1, errors.New("negative span")
	case len(digits) == 0:
		return 1, nil
	case len(regexp.MustCompile(`[^\d]`).FindAllString(digits, -1)) > 0:
		return 0, errors.New("numbers, please")
	}
	for j := 0; j < len(digits)-(span-1); j++ {
		remaining := digits[j:]
		substring := remaining[:span]
		var t int64 = 1
		for _, r := range substring {
			i, _ := strconv.Atoi(string(r))
			t = t * int64(i)
		}
		if t > product {
			product = t
		}
	}
	return product, nil
}
