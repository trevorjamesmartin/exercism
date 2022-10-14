package protein

import (
	"errors"
	"math"
)

var ErrInvalidBase = errors.New("invalid base")
var ErrStop = errors.New("stop")

// true when (a) includes (b), false otherwise
func Includes(a []string, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == b {
			return true
		}
	}
	return false
}

// a sequence of 3 letter strings
func Triplets(s string) []string {
	var result []string
	limit := len(s)
	for i := 0; i < limit; i += 3 {
		triplet := s[i : i+int(math.Min(float64(limit-i), 3))]
		result = append(result, triplet)
	}
	return result
}

func FromRNA(rna string) ([]string, error) {
	var protein []string
	for _, c := range Triplets(rna) {
		acid, e := FromCodon(c)
		switch e {
		case ErrStop:
			return protein, nil
		case ErrInvalidBase:
			return protein, e
		default:
			if len(acid) > 0 && !Includes(protein, acid) {
				protein = append(protein, acid)
			}
		}
	}
	return protein, nil
}

func FromCodon(codon string) (s string, e error) {
	switch codon {
	case "AUG":
		s = "Methionine"
	case "UUU", "UUC":
		s = "Phenylalanine"
	case "UUA", "UUG":
		s = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		s = "Serine"
	case "UAU", "UAC":
		s = "Tyrosine"
	case "UGU", "UGC":
		s = "Cysteine"
	case "UGG":
		s = "Tryptophan"
	case "UAA", "UAG", "UGA":
		s = ""
		e = ErrStop
	default:
		s = ""
		e = ErrInvalidBase
	}
	return s, e
}
