package anagram

import (
	"fmt"
)

// Sum candidate Ascii codes and
// compare with the sum of subject Ascii code

func Detect(subject string, candidates []string) []string {

	// todo: dynamic size array?
	var anagrams []string

	var subjectAsciiCodeSum = 0
	for _, subjectAsciiCode := range subject {
		subjectAsciiCodeSum = subjectAsciiCodeSum + int(subjectAsciiCode)
	}

	for _, candidate := range candidates {
		var candidateAsciiCodeSum = 0
		for _, candidateAsciiCode := range candidate {
			candidateAsciiCodeSum = candidateAsciiCodeSum + int(candidateAsciiCode)
		}

		fmt.Println(candidate, ":", candidateAsciiCodeSum, " | ", subject, ":", subjectAsciiCodeSum)

		if subjectAsciiCodeSum == candidateAsciiCodeSum {
			// add candidate to the array
			anagrams[0] = candidate
		}
	}
	return anagrams
}
