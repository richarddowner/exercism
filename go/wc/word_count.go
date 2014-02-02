/*
Package wc implements WordCount;
a function to count occurances of
words in a string.
*/
package wc

import (
	//"fmt"
	"strings"
)

type Histogram map[string]int

func WordCount(input string) Histogram {

	histogram := make(Histogram)

	words := strings.Split(input, " ")

	for _, word := range words {
		count := 0
		for _, possibleMatch := range words {
			if word == possibleMatch {
				// fmt.Println("Match: ", word, ":", possibleMatch)
				count++
			}
		}
		histogram[word] = count
	}

	return histogram
}
