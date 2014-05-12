package scrabble_score

import (
	"strings"
)

var store = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

func Score(word string) (score int) {
	for _, char := range word {
		for key, value := range store {
			if strings.Contains(key, strings.ToUpper(string(char))) {
				score += value
				break
			}
		}
	}
	return
}
