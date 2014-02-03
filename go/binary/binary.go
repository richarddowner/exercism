package binary

import (
	"math"
	"unicode/utf8"
)

func ToDecimal(binary string) (decimal int) {
	exponent := 0
	for bit := utf8.RuneCountInString(binary) - 1; bit >= 0; bit-- {
		if binary[bit] == '1' {
			// decimal += 2^exponent
			decimal += int(math.Pow(float64(2), float64(exponent)))
		} else if binary[bit] != '0' {
			// not a binary number
			return 0
		}
		exponent++
	}
	return
}
