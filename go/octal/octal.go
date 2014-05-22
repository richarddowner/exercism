// Package octal implements ToDecimal.
package octal

import (
	"math"
)

// ToDecimal takes an octal string.
// Returns the converted decimal integer.
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
