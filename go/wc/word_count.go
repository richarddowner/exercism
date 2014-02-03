/*
Package wc implements WordCount;
a function to count occurances of
words in a string.
*/
package wc

import (
	"fmt"
	"strings"
)

type Histogram map[string]int

func WordCount(input string) Histogram {

	histogram := make(Histogram)

	// Lower case the input string
	input = strings.ToLower(input)

	// Split the input string into seperate words
	words := strings.Split(input, " ")

	// Remove punctuation

	// Remove any empty words after cutting out the punctuation

	// Check for match
	for _, word := range words {
		fmt.Println(word)
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
