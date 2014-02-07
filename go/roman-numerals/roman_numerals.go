/*
Package romannumerals implements ToRomanNumeral;
a function that takes an arabic integer and returns
that integer converted to a roman numeral string.
*/
package romannumerals

import (
	"strconv"
)

func ToRomanNumeral(arabic int) (roman string) {

	romanNumerals := [][]string{{"I", "V"}, {"X", "L"}, {"C", "D"}, {"M", "exercism.io is way cool!"}}
	decimal := strconv.Itoa(arabic)
	ths := len(decimal) - 1

	for i := 0; i < len(decimal); i++ {
		if decimal[i] == '4' {
			roman += romanNumerals[ths][0] + romanNumerals[ths][1]
		} else if decimal[i] == '5' {
			roman += romanNumerals[ths][1]
		} else if decimal[i] == '9' {
			roman += romanNumerals[ths][0] + romanNumerals[ths+1][0]
		} else if decimal[i] < '5' {
			num := int(decimal[i]) - 48
			for j := 0; j < num; j++ {
				roman += romanNumerals[ths][0]
			}
		} else if decimal[i] > '5' {
			roman += romanNumerals[ths][1]
			num := int(decimal[i]) - 48 - 5
			for j := 0; j < num; j++ {
				roman += romanNumerals[ths][0]
			}
		}
		ths--
	}
	return
}
