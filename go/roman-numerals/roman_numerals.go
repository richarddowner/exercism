package romannumerals

import (
	// "fmt"
	"strconv"
)

func ToRomanNumeral(arabic int) (roman string) {

	romanNumerals := [][]string{{"I", "V"}, {"X", "L"}, {"C", "D"}, {"M", ""}}
	buildString := make([]string, 0)

	decimal := strconv.Itoa(arabic)

	th := len(decimal) - 1
	for i := 0; i < len(decimal); i++ {
		if decimal[i] == '4' {
			buildString = append(buildString, romanNumerals[th][0]+romanNumerals[th][1])
		} else if decimal[i] == '5' {
			buildString = append(buildString, romanNumerals[th][1])
		} else if decimal[i] == '9' {
			buildString = append(buildString, romanNumerals[th][0]+romanNumerals[th+1][0])
		} else if decimal[i] < '5' {
			num := int(decimal[i]) - 48
			for j := 0; j < num; j++ {
				buildString = append(buildString, romanNumerals[th][0])
			}
		} else if decimal[i] > '5' {
			buildString = append(buildString, romanNumerals[th][1])
			num := int(decimal[i]) - 48 - 5
			for j := 0; j < num; j++ {
				buildString = append(buildString, romanNumerals[th][0])
			}
		}
		th--
	}
	for _, char := range buildString {
		roman += char
	}
	return
}
