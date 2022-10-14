package sieve

func Sieve(limit int) []int {
	primes := []int{}
	m := map[int][]int{}
	for q := 2; q <= limit; q++ {
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
			if q*q < limit {
				m[q*q] = []int{q}
			}
		}
	}
	return primes
}
