package dna

type DNA struct {
	Strand string
	Count  int
}

type Histogram struct {
}

func NewDNA(strand string) DNA {
	dna := DNA{
		Strand: strand,
	}
	return dna
}
