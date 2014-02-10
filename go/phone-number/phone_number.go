package phonenumber

import (
	"regexp"
	"strings"
)

func FormatPhoneNumber(input string) (output string) {
	invalidNumber := "0000000000"
	output = strings.Join(regexp.MustCompile("([0-9])+").FindAllString(input, -1), "")
	if len(output) < 10 {
		return invalidNumber
	}
	if len(output) == 11 && output[0] == '1' {
		output = output[1:]
	}
	if len(output) == 11 && output[0] != '1' {
		return invalidNumber
	}
	if len(output) > 11 {
		return invalidNumber
	}
	return
}

func Number(input string) (output string) {
	return FormatPhoneNumber(input)
}

func AreaCode(input string) (output string) {
	output = FormatPhoneNumber(input)
	return output[0:3]
}

func Format(input string) (output string) {
	input = FormatPhoneNumber(input)
	return "(" + AreaCode(input) + ") " + input[3:6] + "-" + input[6:10]
}
