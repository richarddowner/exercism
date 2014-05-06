package raindrops

import (
	"strconv"
)

func removeFactor(number int, factor int) (output int) {
	return number - (number / factor)
}

func Convert(number int) (output string) {
	if number%3 == 0 {
		output += "Pling"
		number = removeFactor(number, 3)
	}
	if number%5 == 0 {
		output += "Plang"
		number = removeFactor(number, 5)
	}
	if number%7 == 0 {
		output += "Plong"
		number = removeFactor(number, 7)
	}
	if output == "" {
		return strconv.Itoa(number)
	}
	return
}
