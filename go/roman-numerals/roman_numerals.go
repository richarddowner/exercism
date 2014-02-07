package romannumerals

// I : 1
// V : 5
// X : 10

// X : 10
// L : 50

// C : 100
// D : 500

// M : 1000

import (
	"fmt"
	"strconv"
)

func ToRomanNumeral(arabic int) string {

	roman := make([]string, 0)

	arabic = 545

	romanNumerals := [][]string{{"I", "V"}, {"X", "L"}, {"C", "D"}, {"M", ""}}

	decimal := strconv.Itoa(arabic)

	fmt.Println("decimal = ", decimal)

	th := len(decimal) - 1
	for i := 0; i < len(decimal); i++ {
		// fmt.Println("i = ", i)
		// fmt.Println("decimal[i] = ", decimal[i])
		if decimal[i] == '4' {
			// fmt.Println("th", th)
			// fmt.Println("romanNumerals[count][0] = ", romanNumerals[th][0])
			// fmt.Println("romanNumerals[count][1] = ", romanNumerals[count][1])
			roman = append(roman, romanNumerals[th][0]+romanNumerals[th][1])
		}
		if decimal[i] == '5' {
			// fmt.Println("th", th)
			roman = append(roman, romanNumerals[th][1])
		}
		th--
	}

	fmt.Println("roman = ", roman)

	return "I"
}

// cases

// x = 0
// x = 1 - 3
// x = 4
// x = 5
// x = 6 - 8
// x = 9

// x = 0
// x = 4
// x = 5
// x = 9
// x < 5 && x % 5 = r
// x > 5 && x % 5 = r

// 48
// tenths
// 4 % 5 = 4 ~ if 4: take (X L)
// ones
// 8 % 5 = 3 ~ (V I I I)

// 59
// tenths
// 5 % 5 = 0 ~ if 0: take (L)
// ones
// 9 % 5 = 4 ~ (I X)

// yes
// 141
// hundreths
// 1 % 5 = 1 (C)
// tenths
// 4 % 5 = 4 ~ if 4: take (X L)
// ones
// 1 % 5 = 1 (I)

// no
// 575
// hundreths
// 5 % 5 = 0 ~ if 0: take (D)
// tenths
// 7 % 5 = 2 ~ (X X)
// ones
// 5 % 5 = 0 ~ if 0: take (V)
