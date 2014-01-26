package hamming

import (
	"math"
)

func Distance(strandA string, strandB string) int {

	distance := 0

	for i := 0; i < int(math.Max(float64(len(strandA)), float64(len(strandB)))); i++ {
		if (len(strandA)) > i && (len(strandB)) > i {
			if strandA[i] != strandB[i] {
				distance = distance + 1
			}
		}
	}
	return distance
}
