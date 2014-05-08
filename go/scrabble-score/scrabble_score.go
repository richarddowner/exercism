package scrabble_score

import (
	"fmt"
	// "strconv"
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

	for i := 0; i < 7; i++ {
		fmt.Println(word)
		for j := 0; j < len(word); j++ {
			fmt.Println(strings.ToUpper(string(word[j])))
			fmt.Println(b[0][i])
			fmt.Println(strings.Contains(b[0][i], string(word[j])))

		}
		fmt.Println()
	}

	// for i := 0; i < len(word); i++ {
	// 	for j := 0; j < len(b); j++ {

	// 		fmt.Println(strings.ToUpper(string(word[i])))
	// 		fmt.Println(b[0][j])
	// 		fmt.Println(strings.Contains(b[0][j], string(word[i])))

	// 		fmt.Println()

	// 		if strconv.Itoa(int(word[i])) == b[0][j] {
	// 			var point, _ = strconv.Atoi(b[1][j])
	// 			score += point
	// 		}
	// 	}
	// }
	return
}
