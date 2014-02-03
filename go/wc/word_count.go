/*
Package wc implements WordCount;
a function to count occurances of
words in a string.
*/
package wc

import (
	"regexp"
	"strings"
)

type Histogram map[string]int

func WordCount(input string) Histogram {
	histogram := make(Histogram)

	// Lower case the input string
	input = strings.ToLower(input)

	// Split and remove punctuation with regex
	words := regexp.MustCompile("([a-z0-9])+").FindAllString(input, -1)

	// Check for match
	for _, word := range words {
		matches := 0
		for _, possibleMatch := range words {
			if word == possibleMatch {
				matches++
			}
		}
		histogram[word] = matches
	}
	return histogram
}
