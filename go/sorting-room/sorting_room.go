package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf(`This is the number %.1f`, f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf(`This is a box containing the number %.1f`, float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	var i interface{} = fnb
	_, ok := i.(FancyNumber)
	if !ok {
		return 0
	}
	n, err := strconv.Atoi(fnb.Value())
	if err != nil {
		return 0
	}
	return n
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf(`This is a fancy box containing the number %.1f`, float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	var ni int
	var ok bool
	ni, ok = i.(int)
	if ok {
		return DescribeNumber(float64(ni))
	}
	var nf64 float64
	nf64, ok = i.(float64)
	if ok {
		return DescribeNumber(nf64)
	}
	var nb NumberBox
	nb, ok = i.(NumberBox)
	if ok {
		return DescribeNumberBox(nb)
	}
	var fnb FancyNumberBox
	fnb, ok = i.(FancyNumberBox)
	if ok {
		return DescribeFancyNumberBox(fnb)
	}
	return "Return to sender"
}
