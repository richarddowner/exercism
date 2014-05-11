package scrabble_score

import (
	"strings"
)

var letters = []string{
	"AEIOULNRST",
	"DG",
	"BCMP",
	"FHVWY",
	"K",
	"JX",
	"QZ",
}

var scores = []int{1, 2, 3, 4, 5, 8, 10}

func Score(word string) (score int) {
	for _, char := range word {
		for i := 0; i < 7; i++ {
			if strings.Contains(letters[i], strings.ToUpper(string(char))) {
				score += scores[i]
				break
			}
		}
	}
	return
}
