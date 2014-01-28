package anagram

import (
	"fmt"
)

func Detect(subject string, candidates []string) []string {

	var anagrams = make([]string, 0)

	for _, candidate := range candidates {

		// check for same length
		if len(subject) == len(candidate) {

			var matches = 0
			for i, char := range subject {
				if char == candidate[i] {
					matches++
				}
			}
			if matches == len(subject) {
				// add candidate to the array
			}

		}

	}

	return anagrams
}

// for each candidate
// loop over subject and candidate checking if characters match
// if no match on one of the characters break
// else add the candidate to the anagram list
