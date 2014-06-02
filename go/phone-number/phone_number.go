package phonenumber

import (
	"fmt"
	"regexp"
	"strings"
)

const INVALID_NUMBER = "0000000000"

var numbersRegex *regexp.Regexp = regexp.MustCompile("([0-9])+")

// Number takes a phone number as a string of digits and validates the
// digits as a phone number. Returns an invalid number or valid number.
func Number(digits string) string {
	validDigits := strings.Join(numbersRegex.FindAllString(digits, -1), "")
	length := len(validDigits)
	switch {
	case length < 10:
		return INVALID_NUMBER
	case length == 10:
		return validDigits
	case (length == 11 && validDigits[0] == '1'):
		return validDigits[1:]
	}
	return INVALID_NUMBER
}

// AreaCode takes a phone number as a string of digits and returns the areacode.
func AreaCode(digits string) string {
	return Number(digits)[0:3]
}

// Format takes a phone number as a string of digits, processes the numbers
// and returns them in a correct phone number format.
func Format(digits string) string {
	validDigits := Number(digits)
	return fmt.Sprintf("(%s) %s-%s", validDigits[0:3], validDigits[3:6], validDigits[6:10])
}
