package anagram

import (
	"fmt"
)

// Sum candidate Ascii codes and
// compare with the sum of subject Ascii code

func Detect(subject string, candidates []string) []string {

	// todo: dynamic size array?
	var anagrams = make([]string, 0)

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

		if subjectAsciiCodeSum == candidateAsciiCodeSum && subject != candidate {
			// add candidate to the array
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}
