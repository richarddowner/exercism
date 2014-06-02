package phonenumber

import (
	"fmt"
	"regexp"
	"strings"
)

const INVALID_NUMBER = "0000000000"

var numbersRegex *regexp.Regexp = regexp.MustCompile("([0-9])+")

func Number(input string) string {
	output := strings.Join(numbersRegex.FindAllString(input, -1), "")
	length := len(output)
	switch {
	case length < 10:
		return INVALID_NUMBER
	case length == 10:
		return output
	case (length == 11 && output[0] == '1'):
		return output[1:]
	}
	return INVALID_NUMBER
}

func AreaCode(input string) string {
	return Number(input)[0:3]
}

func Format(input string) string {
	output := Number(input)
	return fmt.Sprintf("(%s) %s-%s", output[0:3], output[3:6], output[6:10])
}
