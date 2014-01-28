package anagram

import (
	"fmt"
)

func Detect(subject string, candidates []string) []string {

	var anagrams = make([]string, 0)

	// for each candidate
	for _, candidate := range candidates {

		// check for same length
		if len(subject) == len(candidate) {

			var matches = 0
			// for each char in subject
			for _, subjectChar := range subject {
				// check each char of candidate
				for _, candidateChar := range candidate {

					if subjectChar == candidateChar {
						matches++
					}
				}
			}
			if matches == len(subject) {
				// add candidate to the array
				fmt.Println("anagram: ", subject, candidate)
				anagrams = append(anagrams, candidate)
			}

		}

	}

	return anagrams
}

// for each candidate
// loop over subject and candidate checking if characters match
// if no match on one of the characters break
// else add the candidate to the anagram list
