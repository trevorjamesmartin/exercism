package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	f, exists := cb[file]
	if !exists {
		return 0
	}
	t := 0
	for _, o := range f {
		if o {
			t++
		}
	}
	return t
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	k := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	t := 0
	if rank < 1 || rank > 8 {
		return 0
	}
	for _, file := range k {
		f := cb[file]
		o := f[rank-1]
		if o {
			t++
		}
	}
	return t
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	k := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	t := 0
	for _, file := range k {
		t += len(cb[file])
	}
	return t
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	k := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	t := 0
	for _, file := range k {
		for _, o := range cb[file] {
			if o {
				t++
			}
		}
	}
	return t
}
