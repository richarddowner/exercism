package dna

type DNA struct {
	Strand string
	Count  int
}

func NewDNA(strand string) DNA {
	dna := DNA{
		Strand: strand,
		Count:  0,
	}
	return dna
}
