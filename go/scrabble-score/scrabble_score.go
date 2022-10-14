package scrabble

import (
	"fmt"
	"strings"
)

type Rule struct {
	key   rune
	value int
	left  *Rule
	right *Rule
}

func (rule *Rule) Value() int {
	return rule.value
}

func (rule *Rule) NewRule(r rune, v int) {
	switch {
	case r > rule.key:
		if rule.right != nil {
			rule.right.NewRule(r, v)
		} else {
			rule.right = &Rule{key: r, value: v}
		}
	case r < rule.key:
		if rule.left != nil {
			rule.left.NewRule(r, v)
		} else {
			rule.left = &Rule{key: r, value: v}
		}
	case r == rule.key && v != rule.value:
		// update value
		rule.value = v
	}
}

func (rule *Rule) Points(key rune) int {
	switch {
	case rule.key == key:
		return rule.value
	case key < rule.key:
		return rule.left.Points(key)
	case key > rule.key:
		return rule.right.Points(key)
	default:
		return 0
	}
}

func (rule *Rule) ForEach(fn func(r rune, v int)) {
	if rule.left != nil {
		rule.left.ForEach(fn)
	}
	if rule.right != nil {
		rule.right.ForEach(fn)
	}
	fn(rule.key, rule.value)
}

func (rule *Rule) Print() {
	if rule.left != nil {
		fmt.Print("\n<")
		rule.left.Print()
	}
	fmt.Printf("%v", rule.value)
	if rule.right != nil {
		fmt.Print("\n>")
		rule.right.Print()
	}
}

func ScrabbleRules() *Rule {
	r := &Rule{key: 'A', value: 1}
	for _, r1 := range []rune{'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'} {
		r.NewRule(r1, 1)
	}
	r.NewRule('D', 2)
	r.NewRule('G', 2)
	for _, r3 := range []rune{'B', 'C', 'M', 'P'} {
		r.NewRule(r3, 3)
	}
	for _, r4 := range []rune{'F', 'H', 'V', 'W', 'Y'} {
		r.NewRule(r4, 4)
	}
	r.NewRule('K', 5)
	r.NewRule('J', 8)
	r.NewRule('X', 8)
	r.NewRule('Q', 10)
	r.NewRule('Z', 10)
	return r
}

func Score(word string) int {
	g := ScrabbleRules()
	g.Print()
	var wordvalue = 0
	for _, letter := range strings.ToUpper(word) {
		wordvalue += g.Points(letter)
	}
	return wordvalue
}
