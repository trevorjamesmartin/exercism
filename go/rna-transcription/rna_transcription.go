package strand

func ToRNA(dna string) string {
	var s []rune
	for _, r := range dna {
		switch r {
		case 'A':
			s = append(s, 'U')
		case 'C':
			s = append(s, 'G')
		case 'T':
			s = append(s, 'A')
		case 'G':
			s = append(s, 'C')
		}
	}
	return string(s)
}
