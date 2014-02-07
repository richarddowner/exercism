package romannumerals

import (
	// "fmt"
	"strconv"
)

func ToRomanNumeral(arabic int) (returnRoman string) {

	romanNumerals := [][]string{{"I", "V"}, {"X", "L"}, {"C", "D"}, {"M", ""}}
	roman := make([]string, 0)

	// arabic = 543 // remove me
	decimal := strconv.Itoa(arabic)

	th := len(decimal) - 1
	for i := 0; i < len(decimal); i++ {
		if decimal[i] == '4' {
			roman = append(roman, romanNumerals[th][0]+romanNumerals[th][1])
		} else if decimal[i] == '5' {
			roman = append(roman, romanNumerals[th][1])
		} else if decimal[i] == '9' {
			roman = append(roman, romanNumerals[th][0]+romanNumerals[th+1][0])
		} else if decimal[i] < '5' {
			num := int(decimal[i]) - 48
			for j := 0; j < num; j++ {
				roman = append(roman, romanNumerals[th][0])
			}
		} else if decimal[i] > '5' {
			roman = append(roman, romanNumerals[th][1])
			num := int(decimal[i]) - 48 - 5
			for j := 0; j < num; j++ {
				roman = append(roman, romanNumerals[th][0])
			}
		}
		th--
	}
	// convert []string to string and return
	for _, s := range roman {
		returnRoman += s
	}
	return returnRoman
}
