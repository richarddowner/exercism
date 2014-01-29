package hamming

func Distance(strandA string, strandB string) (distance int) {
	for i := 0; len(strandA) > i && len(strandB) > i; i++ {
		if strandA[i] != strandB[i] {
			distance++
		}
	}
	return distance
}
