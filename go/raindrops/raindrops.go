package raindrops

import (
	"strconv"
)

func toRainDrops(number int, factor int, word string) (output string) {
	if number%factor == 0 {
		output += word
	}
	return
}

func Convert(number int) (output string) {
	output += toRainDrops(number, 3, "Pling")
	output += toRainDrops(number, 5, "Plang")
	output += toRainDrops(number, 7, "Plong")
	if output == "" {
		return strconv.Itoa(number)
	}
	return
}
