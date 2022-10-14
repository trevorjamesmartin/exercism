package variablelengthquantity

import (
	"errors"
	"math"
)

// integer to binary-notation
func iToBn(a int) []byte {
	b := a / 2
	c := 1
	for b > 0 {
		b >>= 1
		c++
	}
	total := 0
	var arr []byte
	for i := c; i >= 0; i-- {
		power := int(math.Pow(2, float64(i)))
		var bit = 0
		if total+power <= a {
			total += power
			bit = 1
		}
		arr = append(arr, byte(bit))
		c--
	}
	return arr
}

// divide into groups of 7
func g7(a []byte) [][]byte {
	var arr [][]byte
	var c int = 0
	var b []byte
	for i := len(a) - 1; i >= 0; i-- {
		c++
		bit := a[i]
		if c < 8 {
			b = append([]byte{bit}, b...)
		} else {
			arr = append([][]byte{b}, arr...)
			b = []byte{bit}
			c = 1
		}
	}
	if len(b) > 0 {
		zeros := 7 - len(b)
		for zeros > 0 && zeros < 7 {
			zeros--
			b = append([]byte{0}, b...)
		}
		arr = append([][]byte{b}, arr...)
	}
	return arr
}

// binary-notation to integer
func bnToI(group []byte) int {
	gt := 0
	groupEnd := len(group) - 1
	for i, b := range group {
		y := int(math.Pow(2, float64(groupEnd-i)))
		if b == 1 {
			gt += y
		}
	}
	return gt
}

// vlq-encode unsigned 32-bit integer(s)
func EncodeVarint(input []uint32) []byte {
	var encoded []byte
	for _, u := range input {
		groups := g7(iToBn(int(u)))
		lastByte := len(groups) - 1
		for i := 0; i <= lastByte; i++ {
			val := 0
			if i < lastByte {
				val += 128
			}
			gv := bnToI(groups[i])
			val += gv
			switch {
			case i > 0, i == 0 && gv > 0:
				encoded = append(encoded, byte(val))
			case i == 0 && u == 0 && lastByte == 0:
				encoded = append(encoded, byte(0))
			}
		}
	}
	return encoded
}

// vlq-decode byte(s)
func DecodeVarint(input []byte) ([]uint32, error) {
	var sequence []uint32
	var whole []byte
	var pieces int
	for _, b := range input {
		piece := iToBn(int(b))
		start := len(piece) - 8
		if start < 0 {
			zeros := 8 - len(piece)
			for zeros > 0 && zeros < 7 {
				zeros--
				piece = append([]byte{0}, piece...)
			}
			start = 0
		}
		whole = append(whole, piece[start+1:]...)
		if b > 127 {
			pieces++
		} else {
			sequence = append(sequence, uint32(bnToI(whole)))
			pieces = 0
			whole = []byte{}
		}
	}
	if len(sequence) == 0 {
		return sequence, errors.New("incomplete sequence")
	}
	return sequence, nil
}
