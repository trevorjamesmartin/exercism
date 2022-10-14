package romannumerals

import (
	"errors"
	"strconv"
	"strings"
)

type Numeral struct {
	Decimal int
}

func (n *Numeral) Series() []int {
	s := strconv.Itoa(n.Decimal)
	z := len(s) - 1
	var result []int
	for _, c := range s {
		trail := strings.Repeat("0", z)
		num, _ := strconv.Atoi(string(c) + trail)
		result = append(result, num)
		z--
	}
	return result
}

func (n *Numeral) Roman() (s string, e error) {
	defer func() {
		r := recover()
		if r != nil {
			e = r.(error)
			s = ""
		}
	}()

	if n.Decimal <= 0 || n.Decimal > 3999 {
		es := strconv.Itoa(n.Decimal)
		e = errors.New("Roman Numeral out of range " + es)
		return s, e
	}

	for _, c := range n.Series() {
		switch {
		case c == 0:
			s += ""
		case c == 1:
			s += "I"
		case c == 2:
			s += "II"
		case c == 3:
			s += "III"
		case c == 4:
			s += "IV"
		case c == 5:
			s += "V"
		case c == 6:
			s += "VI"
		case c == 7:
			s += "VII"
		case c == 8:
			s += "VIII"
		case c == 9:
			s += "IX"
		case c >= 10 && c < 100:
			switch c {
			case 10:
				s += "X"
			case 20:
				s += "XX"
			case 30:
				s += "XXX"
			case 40:
				s += "XL"
			case 50:
				s += "L"
			case 60:
				s += "LX"
			case 70:
				s += "LXX"
			case 80:
				s += "LXXX"
			case 90:
				s += "XC"
			}
		case c >= 100 && c < 1000:
			switch c {
			case 100:
				s += "C"
			case 200:
				s += "CC"
			case 300:
				s += "CCC"
			case 400:
				s += "CD"
			case 500:
				s += "D"
			case 600:
				s += "DC"
			case 700:
				s += "DCC"
			case 800:
				s += "DCCC"
			case 900:
				s += "CM"
			}
		case c >= 1000:
			switch c {
			case 1000:
				s += "M"
			case 2000:
				s += "MM"
			case 3000:
				s += "MMM"
			}
		}
	}
	return s, e
}

func ToRomanNumeral(input int) (string, error) {
	n := Numeral{Decimal: input}
	return n.Roman()
}
