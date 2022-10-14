package summultiples

type Multiple struct {
	Value int
	Next  *Multiple
}

func (m *Multiple) Insert(i int) {
	switch {
	// i should be unique
	case m.Value == i:
		return
	case m.Value == 0 && m.Next == nil:
		// store i
		m.Value = i
		// create Next
		m.Next = &Multiple{}
		return
	case m.Next != nil:
		m.Next.Insert(i)
	}
}

func (m *Multiple) Sum() int {
	t := m.Value
	i := m
	for i.Next != nil {
		i = i.Next
		t += i.Value
	}
	return t
}

func SumMultiples(limit int, divisors ...int) int {
	m := &Multiple{}
	for _, i := range divisors {
		if i > 0 {
			for j := 1; i*j < limit; j++ {
				m.Insert(i * j) // catalogue the product(s)
			}
		}
	}
	return m.Sum()
}
