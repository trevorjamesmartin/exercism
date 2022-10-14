package reverse

func Reverse(input string) string {
	a := []rune(input)
	var b []rune
	for n := len(a) - 1; n >= 0; n-- {
		b = append(b, a[n])
	}
	c := string(b)
	return c
}
