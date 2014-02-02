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

	input = strings.ToLower(input)

	words := strings.Split(input, " ")

	for _, word := range words {

		// remove punctuation
		word = strings.Replace(word, ":", "", -1)
		word = strings.Replace(word, "!", "", -1)
		word = strings.Replace(word, "&", "", -1)
		word = strings.Replace(word, "@", "", -1)
		word = strings.Replace(word, "$", "", -1)
		word = strings.Replace(word, "%", "", -1)
		word = strings.Replace(word, "^", "", -1)
		fmt.Println(word)

		matches := 0
		for _, possibleMatch := range words {
			if word == possibleMatch {
				// fmt.Println("Match: ", word, ":", possibleMatch)
				matches++
			}
		}
		histogram[word] = matches
	}

	return histogram
}
