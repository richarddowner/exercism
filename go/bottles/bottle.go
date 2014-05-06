package bottles

import (
	"strconv"
)

var bottles int = 99

func Verse(number int) (output string) {
	bottles = number
	if number > 1 {
		output += strconv.Itoa(bottles) + " bottles of beer on the wall, " + strconv.Itoa(bottles) + " bottles of beer.\n"
		bottles--
		if bottles == 1 {
			output += "Take one down and pass it around, 1 bottle of beer on the wall.\n"
		} else {
			output += "Take one down and pass it around, " + strconv.Itoa(bottles) + " bottles of beer on the wall.\n"
		}
		return
	}
	if number == 1 {
		output += "1 bottle of beer on the wall, 1 bottle of beer.\n"
		output += "Take it down and pass it around, no more bottles of beer on the wall.\n"
		return
	} else {
		output += "No more bottles of beer on the wall, no more bottles of beer.\n"
		output += "Go to the store and buy some more, 99 bottles of beer on the wall.\n"
		return
	}
}
