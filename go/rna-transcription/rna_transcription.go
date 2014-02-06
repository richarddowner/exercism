package strand

// C G A T -> G C U A
func ToRna(dna string) (rna string) {
	for _, base := range dna {
		if base == 'C' {
			rna += "G"
		} else if base == 'G' {
			rna += "C"
		} else if base == 'A' {
			rna += "U"
		} else if base == 'T' {
			rna += "A"
		}
	}
	return
}
