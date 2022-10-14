package grains

import "errors"

func Square(number int) (uint64, error) {
	var e error
	var n uint64
	switch {
	case number < 0:
		e = errors.New("a negative square does not exist on this board")
	case number > 64:
		e = errors.New("the chess board has a total of 64 squares")
	case number == 0:
		e = errors.New("divide grain by zero?")
	case number == 1:
		return 1, e
	default:
		n = 1
		for i := 1; i < number; i++ {
			n = n << 1
		}
	}
	return n, e
}

func Total() uint64 {
	sum := uint64(0)
	for i := 1; i <= 64; i++ {
		n, _ := Square(i)
		sum += n
	}
	return sum
}
