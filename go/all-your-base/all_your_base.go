package allyourbase

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	var outputDigits []int
	if inputBase < 2 {
		return outputDigits, fmt.Errorf("input base must be >= 2")
	}
	if outputBase < 2 {
		return outputDigits, fmt.Errorf("output base must be >= 2")
	}
	start := len(inputDigits) - 1
	total := 0
	// convert to base-10
	for i, digit := range inputDigits {
		if digit >= inputBase || digit < 0 {
			return outputDigits, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
		value := float64(digit) * math.Pow(float64(inputBase), float64(start-i))
		total += int(value)
	}
	if outputBase == 10 {
		for _, s := range strings.Split(fmt.Sprintf("%d", total), "") {
			digit, _ := strconv.Atoi(s)
			outputDigits = append(outputDigits, digit)
		}
		return outputDigits, nil
	}

	// determine number of positions in outputBase
	b := total / outputBase
	c := 1
	for a := 1; b > 0; a++ {
		b = b / outputBase
		c++
	}

	last := 0

	// convert from base-10
	for i := c; i > 0; i-- {
		// positional notation multiplier
		x := int(math.Pow(float64(outputBase), float64(i-1)))
		y := 1
		for (y+1)*x+last < total {
			// y * (base ^ position)
			y++
		}
		if y*x+last > total {
			// overflow
			y--
		}
		last += y * x
		outputDigits = append(outputDigits, y)
		if i == 1 {
			offBy := total - last
			for z := len(outputDigits) - 1; z > 0 && offBy > 0; z-- {
				// underflow
				if outputDigits[z]+1 == outputBase {
					outputDigits[z] = 0
				} else {
					outputDigits[z]++
					offBy--
				}
			}
		}
	}
	return outputDigits, nil
}
