package romannumerals

import (
	"fmt"
	"strconv"
)

func ToRomanNumeral(arabic int) string {

	roman := make([]string, 0)

	arabic = 6

	romanNumerals := [][]string{{"I", "V"}, {"X", "L"}, {"C", "D"}, {"M", ""}}

	decimal := strconv.Itoa(arabic)

	fmt.Println("decimal = ", decimal)

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
			// add romanNumerals[th][1]
			roman = append(roman, romanNumerals[th][1])
			// convert to int
			num := int(decimal[i]) - 48
			// loop over adding romanNumerals[th][0] each time
			for j := 0; j < num; j++ {
				roman = append(roman, romanNumerals[th][0])
			}
		}
		th--
	}

	fmt.Println("roman = ", roman)
	return "I"
}
