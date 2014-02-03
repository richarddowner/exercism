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
	hist := Histogram{}

	// Normalize the input string by converting to lower case
	input = strings.ToLower(input)

	// Split out the words and remove punctuation with regex
	words := regexp.MustCompile("([a-z0-9])+").FindAllString(input, -1)

	// Increment occurrences of word
	for _, word := range words {
		hist[word]++
	}
	return hist
}
