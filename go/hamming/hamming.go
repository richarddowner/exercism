/*
Package hamming implements Distance;
a simple function for calculating the
hamming distance between two DNA strands.
*/
package hamming

func Distance(strandA, strandB string) (mutations int) {
	for i := 0; len(strandA) > i && len(strandB) > i; i++ {
		if strandA[i] != strandB[i] {
			mutations++
		}
	}
	return
}
