package phonenumber

import (
	"regexp"
	"strings"
)

const invalidNumber string = "0000000000"

var numbersRegex *regexp.Regexp = regexp.MustCompile("([0-9])+")

func Number(input string) string {
	return formatPhoneNumber(input)
}

func AreaCode(input string) string {
	output := formatPhoneNumber(input)
	return output[0:3]
}

func Format(input string) string {
	output := formatPhoneNumber(input)
	return "(" + AreaCode(output) + ") " + output[3:6] + "-" + output[6:10]
}

func formatPhoneNumber(input string) string {
	output := strings.Join(numbersRegex.FindAllString(input, -1), "")
	length := len(output)
	if length < 10 {
		return invalidNumber
	}
	if length == 11 && output[0] == '1' {
		output = output[1:]
	}
	if length == 11 && output[0] != '1' {
		return invalidNumber
	}
	if length > 11 {
		return invalidNumber
	}
	return output
}
