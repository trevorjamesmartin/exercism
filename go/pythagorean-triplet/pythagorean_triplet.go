package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	var t []Triplet
	for a := min; a < max; a++ {
		for b := a + 1; b < max; b++ {
			for c := b + 1; c <= max; c++ {
				if a*a+b*b == c*c {
					t = append(t, Triplet{a, b, c})
				}
			}
		}
	}
	return t
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	var result []Triplet
	a, b := 1, 2
	c := p - a - b
	for a < b && b < c {
		aSquared := a * a
		for b < c {
			if aSquared+b*b == c*c {
				result = append(result, Triplet{a, b, c})
			}
			b++
			c = p - a - b
		}
		a++
		b = a + 1
		c = p - a - b
	}
	return result
}
