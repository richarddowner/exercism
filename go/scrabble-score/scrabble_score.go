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
	for _, char := range word {
		for j := 0; j < 7; j++ {
			if strings.Contains(b[0][j], strings.ToUpper(string(char))) {
				point, _ := strconv.Atoi(b[1][j])
				score += point
				break
			}
		}
	}
	return
}
