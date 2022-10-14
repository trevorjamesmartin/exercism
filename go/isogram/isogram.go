package isogram

type RuneTree struct {
	key   rune
	count int
	left  *RuneTree
	right *RuneTree
}

// insert [rune] and return the count
func (art *RuneTree) InCount(r rune) int {
	switch {
	case art.key == r:
		art.count++
		return art.count

	case r < art.key:
		if art.left != nil {
			return art.left.InCount(r)
		} else {
			art.left = &RuneTree{key: r, count: 1}
			return 1
		}

	case r > art.key:
		if art.right != nil {
			return art.right.InCount(r)
		} else {
			art.right = &RuneTree{key: r, count: 1}
			return 1
		}

	default:
		return 0
	}
}

func IsIsogram(word string) bool {
	abc := RuneTree{}
	for _, letter := range word {
		if letter > 'Z' {
			letter -= 32
		}
		if letter >= 'A' && letter <= 'Z' && abc.InCount(letter) > 1 {
			return false
		}
	}
	return true
}
