package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

// word-wrap string to approximate square
func SquareString(s string, vertical bool) []string {
	root := math.Sqrt(float64(len(s)))
	col := math.Ceil(root)
	diff := math.Ceil(col - root)
	r, c := int(col-diff), int(col)
	if r*c < len(s) {
		r += 1
	}
	space := 0
	if vertical {
		c, r = r, c
		space = 1
	}
	n := ((r * c) - len(s))
	stringList := strings.Split(s, "")
	var list []string
	idx := 0
	for i := 0; i < r; i++ {
		row := ""
		if r-n <= i {
			for j := 0; j < c-space && idx < len(s); j++ {
				row += stringList[idx]
				idx++
			}
		} else {
			for j := 0; j < c; j++ {
				row += stringList[idx]
				idx++
			}
		}
		list = append(list, row)
	}
	return list
}

func Encode(pt string) string {
	if len(pt) == 0 {
		return ""
	}
	re := regexp.MustCompile(`[a-z0-9]`)
	s1 := strings.Join(re.FindAllString(strings.ToLower(pt), -1), "")
	horizontal := SquareString(s1, false)
	rows := len(horizontal)
	cols := len(horizontal[0])
	s2 := ""
	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			if c < len(horizontal[r]) {
				s2 += strings.Split(horizontal[r], "")[c]
			} else {
				s2 += " "
			}
		}
	}
	vertical := SquareString(s2, true)
	return strings.Join(vertical, " ")
}
