package phonenumber

import (
	"regexp"
	"strings"
)

func FormatPhoneNumber(input string) string {
	output := strings.Join(regexp.MustCompile("([0-9])+").FindAllString(input, -1), "")
	invalidNumber := "0000000000"
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
	return output
}

func Number(input string) string {
	return FormatPhoneNumber(input)
}

func AreaCode(input string) string {
	output := FormatPhoneNumber(input)
	return output[0:3]
}

func Format(input string) string {
	output := FormatPhoneNumber(input)
	return "(" + AreaCode(output) + ") " + output[3:6] + "-" + output[6:10]
}
