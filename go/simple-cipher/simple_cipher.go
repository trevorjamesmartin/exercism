package cipher

import (
	"regexp"
	"strings"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
type shift struct {
	distance int
}

type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return &shift{distance: 3}
}

func NewShift(distance int) Cipher {
	switch {
	case distance < -25, distance == 0, distance > 25:
		return nil
	default:
		return &shift{distance: distance}
	}
}

func CleanLowerCase(s string) string {
	return strings.Join(regexp.MustCompile(`[a-z]`).FindAllString(strings.ToLower(s), -1), "")
}

func (c shift) Encode(input string) string {
	s := ""
	for _, r := range CleanLowerCase(input) {
		switch {
		case r > 96 && r < 123:
			r2 := int(r) - 97
			r2 = r2 + c.distance
			if r2 < 0 {
				r2 += 26
			} else {
				r2 = r2 % 26
			}
			s += string(rune(r2 + 97))
		}
	}
	return s
}

func (c shift) Decode(input string) string {
	s := ""
	for _, r := range input {
		switch {
		case r > 96 && r < 123:
			r2 := int(r) - 97
			r2 -= c.distance
			if r2 < 0 {
				r2 += 26
			} else {
				r2 = r2 % 26
			}
			s += string(rune(r2 + 97))
		case r > 64 && r < 91:
			r2 := int(r) - 65
			r2 -= c.distance
			if r2 < 0 {
				r2 += 26
			} else {
				r2 = r2 % 26
			}
			s += string(rune(r2 + 65))
		}
	}
	return s
}

func NewVigenere(key string) Cipher {
	re := regexp.MustCompile(`[a-z]`)
	k := strings.Join(re.FindAllString(key, -1), "")
	switch {
	case len(key) == 0:
		return nil
	case len(k) < len(key):
		return nil
	}
	total := 0
	for _, r := range k {
		total += int(r) - 97
	}
	if total == 0 {
		return nil
	}
	return &vigenere{key: key}
}

func (v vigenere) Encode(input string) string {
	i := 0
	key := []rune(v.key)
	s := []rune{}
	for _, r := range CleanLowerCase(input) {
		c := int(r) - 97
		d := int(key[i]) - 97
		e := (c + d) % 26
		s = append(s, rune(e+97))
		i = (i + 1) % len(key)
	}
	return string(s)
}

func (v vigenere) Decode(input string) string {
	i := 0
	key := []rune(v.key)
	s := []rune{}
	for _, r := range CleanLowerCase(input) {
		c := int(r) - 97
		d := int(key[i]) - 97
		e := c - d
		if e < 0 {
			e += 26
		} else {
			e = e % 26
		}
		s = append(s, rune(e+97))
		i = (i + 1) % len(key)
	}
	return string(s)
}
