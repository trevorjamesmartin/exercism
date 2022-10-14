package prime

import "errors"

// Nth returns the nth prime number.
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("index out of range")
	}
	m := map[int][]int{}
	primes := []int{}
	q := 2
	for len(primes) < n {
		if m[q] != nil {
			for i := 0; i < len(m[q]); i++ {
				p := m[q][i]
				if m[p+q] != nil {
					m[p+q] = append(m[p+q], p)
				} else {
					m[p+q] = []int{p}
				}
			}
			m[q] = nil
		} else {
			primes = append(primes, q)
			m[q*q] = []int{q}
		}
		q++
	}
	return primes[n-1], nil
}
