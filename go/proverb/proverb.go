// For want of a horseshoe nail, a kingdom was lost, or so the saying goes.
package proverb

import "fmt"

// generates a proverb using the provided list of strings
func Proverb(rhyme []string) []string {
	var s []string
	if len(rhyme) < 1 {
		return s
	}
	for i := 1; i < len(rhyme); i++ {
		a := rhyme[i-1]
		b := rhyme[i]
		s = append(s, fmt.Sprintf("For want of a %v the %v was lost.", a, b))
	}
	s = append(s, fmt.Sprintf("And all for the want of a %v.", rhyme[0]))
	return s
}
