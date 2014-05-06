package raindrops

import (
	"strconv"
)

var number int = 0

func toRainDrops(number int, factor int, word string) (output string) {
	if number%factor == 0 {
		output += word
		number = number - (number / factor)
	}
	return
}

func Convert(input int) (output string) {
	number = input
	output += toRainDrops(number, 3, "Pling")
	output += toRainDrops(number, 5, "Plang")
	output += toRainDrops(number, 7, "Plong")
	if output == "" {
		return strconv.Itoa(number)
	}
	return
}
