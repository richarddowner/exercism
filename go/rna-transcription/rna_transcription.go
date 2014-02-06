/*
Package strand implements ToRna;
a function to transcribe DNA to RNA.
This is the second step in synthesising
protiens from DNA.

Each base in the DNA strand is transcribed
as follows: C G A T -> G C U A
*/
package strand

func ToRna(dna string) (rna string) {
	for _, base := range dna {
		switch base {
		case 'C':
			rna += "G"
		case 'G':
			rna += "C"
		case 'A':
			rna += "U"
		case 'T':
			rna += "A"
		}
	}
	return
}
