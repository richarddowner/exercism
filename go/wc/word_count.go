/*
Package wc implements WordCount;
a function to count occurances of
words in a string.
*/
package wc

import (
	"fmt"
	"regexp"
	"strings"
)

type Histogram map[string]int

func WordCount(input string) Histogram {

	// Lower case the input string
	input = strings.ToLower(input)

	// Split and remove punctuation in 1 step
	words := regexp.MustCompile("([a-z])+").FindAllString(input, -1)

	fmt.Println(words)

	histogram := make(Histogram)

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
