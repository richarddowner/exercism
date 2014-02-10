package phonenumber

import (
	//"fmt"
	"regexp"
	"strings"
)

func Number(input string) (output string) {
	// remove invalid chars
	output = strings.Join(regexp.MustCompile("([0-9])+").FindAllString(input, -1), "")
	// if the phone number is less than 10 digits assume that it is bad number
	if len(output) < 10 {
		return "0000000000"
	}
	// if the phone number is 11 digits and the first number is 1, trim the 1 and use the first 10 digits
	if len(output) == 11 && output[0] == '1' {
		output = output[1:]
	}
	// if the phone number is 11 digits and the first number is not 1, then it is a bad number
	if len(output) == 11 && output[0] != '1' {
		return "0000000000"
	}
	// if the phone number is more than 11 digits assume that it is a bad number
	if len(output) > 11 {
		return "0000000000"
	}
	return
}

func AreaCode(input string) (output string) {
	return
}

func Format(input string) (output string) {
	return
}
