/*
Package octal implements ToDecimal;
a function that takes an octal string
and returns the converted decimal integer.
*/
package octal

import (
	"math"
)

func ToDecimal(octal string) (decimal int64) {
	exponent := 0
	for i := len(octal) - 1; i >= 0; i-- {
		num := octal[i] - 48
		if num > 9 {
			return 0
		}
		decimal += int64(num) * int64(math.Pow(2, float64(exponent)))
		exponent += 3
	}
	return
}
