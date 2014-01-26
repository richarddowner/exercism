package anagram

import (
	"fmt"
)

func Detect(subject string, candidates []string) []string {

	// todo: dynamic size array?
	var anagrams []string

	var subjectAsciiTotal = 0
	for _, subjectAsciiCode := range subject {
		subjectAsciiTotal = subjectAsciiTotal + int(subjectAsciiCode)
	}

	for _, candidate := range candidates {
		var candidateAsciiTotal = 0
		for _, candidateAsciiCode := range candidate {
			candidateAsciiTotal = candidateAsciiTotal + int(candidateAsciiCode)
		}

		fmt.Println(candidate, ":", candidateAsciiTotal, " | ", subject, ":", subjectAsciiTotal)

		if subjectAsciiTotal == candidateAsciiTotal {
			// add candidate to the array
			anagrams[0] = candidate
		}
	}
	return anagrams
}
