package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var i2 Ints
	for _, n := range i {
		if filter(n) {
			i2 = append(i2, n)
		}
	}
	return i2
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var i2 Ints
	for _, n := range i {
		if !filter(n) {
			i2 = append(i2, n)
		}
	}
	return i2
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var l2 Lists
	for _, c := range l {
		if filter(c) {
			l2 = append(l2, c)
		}
	}
	return l2
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var s2 Strings
	for _, t := range s {
		if filter(t) {
			s2 = append(s2, t)
		}
	}
	return s2
}
