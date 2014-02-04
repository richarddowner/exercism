package allergies

import (
	"strconv"
)

var allergiesList = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(score int) []string {
	allergyMatches := make([]string, 0)
	// convert score to binary byte format
	byte := strconv.FormatInt(int64(score), 2)
	position := 0
	// loop of the byte string backwards
	for bit := len(byte) - 1; bit >= 0; bit-- {
		// if bit is 1 and
		// current position is less than the total length of allergiesList array
		if byte[bit] == '1' && position < len(allergiesList) {
			allergyMatches = append(allergyMatches, allergiesList[position])
		}
		position++
	}
	return allergyMatches
}

func AllergicTo(score int, name string) (fact bool) {
	byte := strconv.FormatInt(int64(score), 2)
	position := 0
	for bit := len(byte) - 1; bit >= 0; bit-- {
		if byte[bit] == '1' && allergiesList[position] == name {
			return true
		}
		position++
	}
	return false
}
