package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (h Histogram, e error) {
	h = Histogram{'A': 0, 'C': 0, 'T': 0, 'G': 0}
	for _, r := range d {
		switch r {
		case 'A':
			h['A']++
		case 'C':
			h['C']++
		case 'G':
			h['G']++
		case 'T':
			h['T']++
		default:
			e = fmt.Errorf("invalid nucleotide: %v", r)
			return h, e
		}
	}
	return h, e
}
