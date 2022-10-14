// Two-fer or 2-fer is short for two for one.
package twofer

import "fmt"

// ShareWith, when given a name, will return a string with the message,
//   "One for [name], one for me."
// if [name] is missing or blank, "you" will replace [name]
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf(`One for %s, one for me.`, name)
}
