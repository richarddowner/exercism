package allergies

import (
	// "fmt"
	"strconv"
)

var commonAllergies = []string{
	"eggs",         // 1
	"peanuts",      // 2
	"shellfish",    // 4
	"strawberries", // 8
	"tomatoes",     // 16
	"chocolate",    // 32
	"pollen",       // 64
	"cats",         // 128
}

func Allergies(score int) []string {
	names := make([]string, 0)

	// convert score to binary byte format
	byte := strconv.FormatInt(int64(score), 2)

	position := 0
	// loop of the byte string backwards
	for bit := len(byte) - 1; bit >= 0; bit-- {
		// if bit is 1 and
		// current position is less than the total length of commonAllergies array
		if byte[bit] == '1' && position < len(commonAllergies) {
			names = append(names, commonAllergies[position])
		}
		position++
	}
	return names
}

/*
Tests
{false, 0, "peanuts"},
{false, 0, "cats"},
{false, 0, "strawberries"},
{true, 1, "eggs"},
{true, 5, "eggs"},
*/

func AllergicTo(score int, name string) (fact bool) {
	return
}
