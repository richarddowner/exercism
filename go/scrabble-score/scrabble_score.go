package scrabble_score

import (
	"strconv"
	"strings"
)

var b = [7][7]string{
	{
		"AEIOULNRST",
		"DG",
		"BCMP",
		"FHVWY",
		"K",
		"JX",
		"QZ",
	},
	{
		"1", "2", "3", "4", "5", "8", "10",
	},
}

func Score(word string) (score int) {
	for i := 0; i < len(word); i++ {
		for j := 0; j < len(b); j++ {
			if strings.ToLower(word[i]) == strings.ToLower(b[0][j]) {
				var point, _ = strconv.Atoi(b[1][j])
				score += point
			}
		}
	}
	return
}
